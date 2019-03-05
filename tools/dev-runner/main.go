package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"golang.org/x/sync/errgroup"

	"powerssl.io/internal/pkg/util"
)

var components = map[string]string{
	"auth":      "serve --config configs/auth/config.yaml",
	"apiserver": "serve --config configs/api/config.yaml",
	//"controller": "serve --config configs/controller/config.yaml",
	//"signer":     "serve --config configs/signer/config.yaml",
	"webapp": "serve --config configs/webapp/config.yaml",
}

func main() {
	logger := util.NewLogger(os.Stdout)

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return util.InterruptHandler(ctx, logger)
	})

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Log("err", err)
	}
	defer watcher.Close()

	of := NewOutlet()

	{
		var padding int = 10 // len(dev-runner)
		for c := range components {
			if l := len(c); l > padding {
				padding = l
			}
		}
		of.Padding = padding
	}

	g.Go(func() error {
		return runComponent(ctx, watcher, of, "vault", "vault", "server -config configs/vault/config.hcl", 0)
	})

	var i int = 1 // skip vault
	for c, a := range components {
		component, arg, idx := c, a, i
		bin := filepath.Join("bin", fmt.Sprintf("powerssl-%s", component))
		g.Go(func() error {
			if err := watcher.Add(bin); err != nil {
				return err
			}
			return runComponent(ctx, watcher, of, component, bin, arg, idx)
		})
		i++
	}

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			logger.Log("err", err)
		}
	}
}

func runComponent(ctx context.Context, watcher *fsnotify.Watcher, of *Outlet, component, bin, arg string, idx int) error {
	cmd := exec.Command(bin, strings.Fields(arg)...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	pipeWait := new(sync.WaitGroup)
	pipeWait.Add(2)
	go of.LineReader(pipeWait, component, idx, stdout, false)
	go of.LineReader(pipeWait, component, idx, stderr, true)

	of.SystemOutput(fmt.Sprintf("starting %s", component))

	finished := make(chan struct{})
	if err := cmd.Start(); err != nil {
		of.SystemOutput(fmt.Sprintf("Failed to start %s: %s", component, err))
		return err
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
			runComponent(ctx, watcher, of, component, bin, arg, idx)
		case <-ctx.Done():
			of.SystemOutput(fmt.Sprintf("Killing %s", component))
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
				if event.Name != bin {
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

	return err
}
