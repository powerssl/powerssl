package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"powerssl.dev/tools/dev-runner/internal"
	"powerssl.dev/tools/dev-runner/internal/component"
)

func makeCmd(comp component.Component, idx int, of *internal.Outlet) (*exec.Cmd, *sync.WaitGroup, error) {
	cmd := exec.Command(comp.Command, strings.Fields(comp.Args)...)
	cmd.Env = append(os.Environ(), comp.Env.Environ()...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}
	pipeWait := new(sync.WaitGroup)
	pipeWait.Add(2)
	go of.LineReader(pipeWait, comp.String(), idx, stdout, false)
	go of.LineReader(pipeWait, comp.String(), idx, stderr, true)

	return cmd, pipeWait, nil
}

func observeComponent(ctx context.Context, of *internal.Outlet, comp component.Component, idx int, interrupt chan struct{}) error {
	var (
		cmd         *exec.Cmd
		finished    chan struct{}
		interrupted bool
		killed      bool
		pipeWait    *sync.WaitGroup
	)

	start := func() error {
		var err error
		if cmd, pipeWait, err = makeCmd(comp, idx, of); err != nil {
			return err
		}
		finished = make(chan struct{})
		interrupted = false
		killed = false

		of.SystemOutput(fmt.Sprintf("Starting %s", comp.String()))
		if err := cmd.Start(); err != nil {
			return fmt.Errorf("failed to start %s: %s", comp.String(), err)
		}

		go func() {
			defer close(finished)
			pipeWait.Wait()
			_ = cmd.Wait()
		}()

		return nil
	}
	if err := start(); err != nil {
		return err
	}

	for {
		select {
		case <-interrupt:
			if interrupted || cmd.Process == nil {
				break
			}
			of.SystemOutput(fmt.Sprintf("Interrupting %s", comp.String()))
			if err := cmd.Process.Signal(os.Interrupt); err != nil {
				of.SystemOutput(fmt.Sprintf("Failed to signal %s", comp.String()))
				return err
			}
			interrupted = true
		case <-ctx.Done():
			if killed || cmd.Process == nil {
				break
			}
			of.SystemOutput(fmt.Sprintf("Killing %s", comp.String()))
			if err := cmd.Process.Kill(); err != nil {
				of.SystemOutput(fmt.Sprintf("Failed to kill %s", comp.String()))
				return err
			}
			killed = true
		case <-finished:
			if killed {
				of.SystemOutput(fmt.Sprintf("Killed %s", comp.String()))
				return nil
			}
			time.Sleep(time.Millisecond * 500)
			if err := start(); err != nil {
				return err
			}
		}
	}
}
