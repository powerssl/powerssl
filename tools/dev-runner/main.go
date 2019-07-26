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
		padding := len("dev-runner")
		for _, c := range component.Components {
			if l := len(c.String()); l > padding {
				padding = l
			}
		}
		of.Padding = padding
	}

	var watcher *fsnotify.Watcher
	{
		var err error
		if watcher, err = fsnotify.NewWatcher(); err != nil {
			of.ErrorOutput(fmt.Sprintf("watcher error: %s", err))
		}
		defer watcher.Close()
	}

	interrupts := make(map[string]chan struct{}, len(component.Components)+1)

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case event, ok := <-watcher.Events:
				if !ok {
					return nil
				}
				// if event.Op&fsnotify.Write != fsnotify.Write {
				// 	break
				// }
				c, ok := interrupts[event.Name]
				if ok {
					c <- struct{}{}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return nil
				}
				return err
			}
		}
	})

	var addComponent func(component.Component)
	{
		var idx int
		addComponent = func(comp component.Component) {
			i := idx
			interrupts[comp.Command] = make(chan struct{})
			g.Go(func() error {
				return observeComponent(ctx, of, comp, i, interrupts[comp.Command])
			})
			idx++
		}
	}

	addComponent(component.Component{
		Command: "vault",
		Args:    "server -config configs/vault/config.hcl",
	})

	for _, comp := range component.Components {
		if err := watcher.Add(comp.Command); err != nil {
			of.ErrorOutput(fmt.Sprintf("watcher error: %s", err))
		}
		addComponent(comp)
	}

	g.Go(func() error {
		return handleVault(of)
	})

	if err := g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			of.ErrorOutput(fmt.Sprintf("error: %s", err))
		}
	}
}

func handleVault(of *internal.Outlet) error {
	var command, args string
	if _, err := os.Stat("local/vault/secret.yaml"); os.IsNotExist(err) {
		command = "bin/powerutil"
		args = "vault --ca local/certs/ca.pem --ca-key local/certs/ca-key.pem"
	} else {
		d, err := ioutil.ReadFile("local/vault/secret.yaml")
		if err != nil {
			return fmt.Errorf("config error: %s", err)
		}
		var config map[string]interface{}
		if err := yaml.Unmarshal(d, &config); err != nil {
			return fmt.Errorf("config error: %s", err)
		}

		command = "vault"
		args = fmt.Sprintf("operator unseal -address https://localhost:8200 -ca-cert local/certs/ca.pem %s", config["keys"].([]interface{})[0].(string))
	}

	comp := component.Component{
		Command: command,
		Args:    args,
	}
	cmd, _, err := makeCmd(comp, 0, of)
	if err != nil {
		return err
	}

	time.Sleep(time.Second)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Failed to start %s: %s", comp.Command, err)
	}

	return nil
}
