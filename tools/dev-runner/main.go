package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"golang.org/x/sync/errgroup"

	"powerssl.io/internal/pkg/util"
	"powerssl.io/tools/dev-runner/internal"
)

var components = map[string]string{
	"auth":       "serve --config configs/auth/config.yaml",
	"apiserver":  "serve --config configs/api/config.yaml",
	"controller": "serve --config configs/controller/config.yaml",
	"signer":     "serve --config configs/signer/config.yaml",
	"webapp":     "serve --config configs/webapp/config.yaml",
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		logger := util.NewLogger(ioutil.Discard)
		return util.InterruptHandler(ctx, logger)
	})

	of := internal.NewOutlet()

	{
		var padding int = 10 // len(dev-runner)
		for c := range components {
			if l := len(c); l > padding {
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
		runComponent(ctx, watcher, of, "vault", "vault", "server -config configs/vault/config.hcl", 0)
	}()

	var i int = 1 // skip vault
	for c, a := range components {
		component, arg, idx := c, a, i
		bin := filepath.Join("bin", fmt.Sprintf("powerssl-%s", component))
		go func() {
			if err := watcher.Add(bin); err != nil {
				of.ErrorOutput(fmt.Sprintf("watcher error: %s", err))
			}
			runComponent(ctx, watcher, of, component, bin, arg, idx)
		}()
		i++
	}

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			of.ErrorOutput(fmt.Sprintf("error: %s", err))
		}
	}
}

func runComponent(ctx context.Context, watcher *fsnotify.Watcher, of *internal.Outlet, component, bin, arg string, idx int) {
	cmd := exec.Command(bin, strings.Fields(arg)...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		of.ErrorOutput(fmt.Sprintf("error: %s", err))
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		of.ErrorOutput(fmt.Sprintf("error: %s", err))
	}

	pipeWait := new(sync.WaitGroup)
	pipeWait.Add(2)
	go of.LineReader(pipeWait, component, idx, stdout, false)
	go of.LineReader(pipeWait, component, idx, stderr, true)

	of.SystemOutput(fmt.Sprintf("starting %s", component))

	finished := make(chan struct{})
	if err := cmd.Start(); err != nil {
		of.ErrorOutput(fmt.Sprintf("Failed to start %s: %s", component, err))
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
}
