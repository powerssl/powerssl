package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/ghodss/yaml"
	"golang.org/x/sync/errgroup"

	"powerssl.io/internal/pkg/util"
	"powerssl.io/tools/dev-runner/internal"
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
		for _, c := range components {
			if l := len(c.Name()); l > padding {
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
		comp := component{
			command: "vault",
			args:    "server -config configs/vault/config.hcl",
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

		comp := component{
			command: command,
			args:    args,
		}
		cmd, pipeWait := comp.Command(of, 0)

		finished := make(chan struct{})
		if err := cmd.Start(); err != nil {
			of.ErrorOutput(fmt.Sprintf("Failed to start %s: %s", comp.command, err))
		}

		go func() {
			defer close(finished)
			pipeWait.Wait()
			cmd.Wait()
		}()

	}()

	for i, c := range components {
		comp := c
		idx := i + 1
		go func() {
			if err := watcher.Add(comp.command); err != nil {
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

func runComponent(ctx context.Context, watcher *fsnotify.Watcher, of *internal.Outlet, comp component, idx int) {
	cmd, pipeWait := comp.Command(of, idx)

	of.SystemOutput(fmt.Sprintf("starting %s", comp.command))

	finished := make(chan struct{})
	if err := cmd.Start(); err != nil {
		of.ErrorOutput(fmt.Sprintf("Failed to start %s: %s", comp.command, err))
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
			of.SystemOutput(fmt.Sprintf("Killing %s", comp.command))
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
				if event.Name != comp.command {
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
