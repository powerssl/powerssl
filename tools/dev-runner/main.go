package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/ghodss/yaml"
	"golang.org/x/sync/errgroup"

	"powerssl.io/powerssl/internal/pkg/component"
	"powerssl.io/powerssl/internal/pkg/util"
	"powerssl.io/powerssl/tools/dev-runner/internal"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		logger := util.NewLogger(ioutil.Discard)
		return util.InterruptHandler(ctx, logger)
	})

	of := internal.NewOutlet()

	{
		var padding int = 10 // len(dev-runner)
		for _, c := range component.Components {
			if l := len(c.String()); l > padding {
				padding = l
			}
		}
		of.Padding = padding
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		of.ErrorOutput(fmt.Sprintf("watcher error: %s", err))
	}
	defer watcher.Close()

	go func() {
		comp := component.Component{
			Command: "vault",
			Args:    "server -config configs/vault/config.hcl",
		}
		runComponent(ctx, watcher, of, comp, 0)
	}()

	go func() {
		var command, args string
		if _, err := os.Stat("local/vault/secret.yaml"); os.IsNotExist(err) {
			command = "powerutil"
			args = "vault --ca local/certs/ca.pem --ca-key local/certs/ca-key.pem"
		} else {
			d, err := ioutil.ReadFile("local/vault/secret.yaml")
			if err != nil {
				of.ErrorOutput(fmt.Sprintf("config error: %s", err))
			}
			var config map[string]interface{}
			if err := yaml.Unmarshal(d, &config); err != nil {
				of.ErrorOutput(fmt.Sprintf("config error: %s", err))
			}

			command = "vault"
			args = fmt.Sprintf("operator unseal -address https://localhost:8200 -ca-cert local/certs/ca.pem %s", config["keys"].([]interface{})[0].(string))
		}
		time.Sleep(time.Second)

		comp := component.Component{
			Command: command,
			Args:    args,
		}
		cmd, pipeWait := makeCmd(comp, comp.String(), 0, of)

		finished := make(chan struct{})
		if err := cmd.Start(); err != nil {
			of.ErrorOutput(fmt.Sprintf("Failed to start %s: %s", comp.Command, err))
		}

		go func() {
			defer close(finished)
			pipeWait.Wait()
			cmd.Wait()
		}()
	}()

	for i, c := range component.Components {
		comp := c
		idx := i + 1
		go func() {
			if err := watcher.Add(comp.Command); err != nil {
				of.ErrorOutput(fmt.Sprintf("watcher error: %s", err))
			}
			runComponent(ctx, watcher, of, comp, idx)
		}()
	}

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			of.ErrorOutput(fmt.Sprintf("error: %s", err))
		}
	}
}

func makeCmd(comp component.Component, name string, idx int, of *internal.Outlet) (*exec.Cmd, *sync.WaitGroup) {
	cmd := exec.Command(comp.Command, strings.Fields(comp.Args)...)
	cmd.Env = append(os.Environ(), comp.Env.Environ()...)

	mustPipe := func(pr io.ReadCloser, err error) io.ReadCloser {
		if err != nil {
			of.ErrorOutput(fmt.Sprintf("error: %s", err))
		}
		return pr
	}

	stdout := mustPipe(cmd.StdoutPipe())
	stderr := mustPipe(cmd.StderrPipe())
	pipeWait := new(sync.WaitGroup)
	pipeWait.Add(2)
	go of.LineReader(pipeWait, name, 0, stdout, false)
	go of.LineReader(pipeWait, name, 0, stderr, true)

	return cmd, pipeWait
}

func runComponent(ctx context.Context, watcher *fsnotify.Watcher, of *internal.Outlet, comp component.Component, idx int) {
	cmd, pipeWait := makeCmd(comp, comp.String(), idx, of)

	of.SystemOutput(fmt.Sprintf("starting %s", comp.Command))

	finished := make(chan struct{})
	if err := cmd.Start(); err != nil {
		of.ErrorOutput(fmt.Sprintf("Failed to start %s: %s", comp.Command, err))
	}

	go func() {
		defer close(finished)
		pipeWait.Wait()
		cmd.Wait()
	}()

	go func() {
		select {
		case <-finished:
			time.Sleep(time.Second)
			runComponent(ctx, watcher, of, comp, idx)
		case <-ctx.Done():
			of.SystemOutput(fmt.Sprintf("Killing %s", comp.Command))
			cmd.Process.Kill()
		}
	}()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Name != comp.Command {
					break
				}
				cmd.Process.Signal(os.Interrupt)
				return
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				of.ErrorOutput(fmt.Sprintf("watcher error: %s", err))
			}
		}
	}()
}
