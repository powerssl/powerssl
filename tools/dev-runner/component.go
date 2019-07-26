package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"powerssl.io/powerssl/internal/pkg/component"
	"powerssl.io/powerssl/tools/dev-runner/internal"
)

func makeCmd(comp component.Component, idx int, of *internal.Outlet) (*exec.Cmd, *sync.WaitGroup) {
	cmd := exec.Command(comp.Command, strings.Fields(comp.Args)...)
	cmd.Env = append(os.Environ(), comp.Env.Environ()...)

	mustPipe := func(pr io.ReadCloser, err error) io.ReadCloser {
		if err != nil {
			of.ErrorOutput(fmt.Sprintf("error: %s", err))
		}
		return pr
	}

	pipeWait := new(sync.WaitGroup)
	pipeWait.Add(2)
	go of.LineReader(pipeWait, comp.String(), idx, mustPipe(cmd.StdoutPipe()), false)
	go of.LineReader(pipeWait, comp.String(), idx, mustPipe(cmd.StderrPipe()), true)

	return cmd, pipeWait
}

func observeComponent(ctx context.Context, of *internal.Outlet, comp component.Component, idx int, interrupt chan struct{}) {
	var (
		cmd      *exec.Cmd
		finished chan struct{}
		pipeWait *sync.WaitGroup
	)

	start := func() {
		cmd, pipeWait = makeCmd(comp, idx, of)
		finished = make(chan struct{})
		go startComponent(of, comp, cmd, pipeWait, finished)
	}
	start()

	go func() {
		for {
			select {
			case <-interrupt:
				if cmd.Process == nil {
					break
				}
				of.SystemOutput(fmt.Sprintf("Interrupting %s", comp.Command))
				cmd.Process.Signal(os.Interrupt)
			case <-ctx.Done():
				if cmd.Process == nil {
					return
				}
				of.SystemOutput(fmt.Sprintf("Killing %s", comp.Command))
				cmd.Process.Kill()
				return
			case <-finished:
				time.Sleep(time.Second / 10)
				start()
			}
		}
	}()
}

func startComponent(of *internal.Outlet, comp component.Component, cmd *exec.Cmd, pipeWait *sync.WaitGroup, finished chan struct{}) {
	of.SystemOutput(fmt.Sprintf("Starting %s", comp.Command))
	if err := cmd.Start(); err != nil {
		of.ErrorOutput(fmt.Sprintf("Failed to start %s: %s", comp.Command, err))
	}
	defer close(finished)
	pipeWait.Wait()
	cmd.Wait()
}
