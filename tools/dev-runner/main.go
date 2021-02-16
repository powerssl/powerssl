package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/ghodss/yaml"
	_ "github.com/improbable-eng/grpc-web/go/grpcweb"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/common/util"
	"powerssl.dev/tools/dev-runner/internal"
	"powerssl.dev/tools/dev-runner/internal/component"
)

func main() {
	var g *errgroup.Group
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx = errgroup.WithContext(ctx)
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
		defer errWrapCloser(watcher, &err)
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
				var c chan struct{}
				c, ok = interrupts[event.Name]
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
			localComp := comp
			g.Go(func() error {
				if val, ok := localComp.Env["POWERSSL_AUTH_TOKEN"]; ok && val == "{{GENERATE}}" {
					var err error
					if localComp.Env["POWERSSL_AUTH_TOKEN"], err = serviceToken(); err != nil {
						of.SystemOutput(err.Error())
						cancel()
					}
				}
				return observeComponent(ctx, of, localComp, i, interrupts[comp.Command])
			})
			idx++
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		cancel()
		of.ErrorOutput(err.Error())
	}
	addComponent(component.Component{
		Name:    "postgres",
		Command: "docker",
		Args:    fmt.Sprintf("run --rm -e POSTGRES_PASSWORD=powerssl -e POSTGRES_DB=powerssl -e POSTGRES_USER=powerssl -e PGDATA=/var/lib/postgresql/data/pgdata -p 5432:5432 -v %s/local/postgresql/data:/var/lib/postgresql/data postgres:13.1", wd),
	})

	if err = internal.WaitForService("localhost:5432", time.Minute); err != nil {
		of.SystemOutput(err.Error())
		cancel()
	}

	if err = handlePostgres(of); err != nil {
		cancel()
		of.ErrorOutput(err.Error())
	}

	addComponent(component.Component{
		Command: "vault",
		Args:    "server -config configs/vault/config.hcl",
	})

	if err = internal.WaitForService("localhost:8200", time.Minute); err != nil {
		cancel()
		of.ErrorOutput(err.Error())
	}

	if err = handleVault(of); err != nil {
		cancel()
		of.ErrorOutput(err.Error())
	}

	for _, comp := range component.Components {
		if comp.Command != "bin/powerssl-temporalserver" {
			continue
		}
		if err = watcher.Add(comp.Command); err != nil {
			of.ErrorOutput(fmt.Sprintf("watcher error: %s", err))
		}
		addComponent(comp)
	}

	if err = internal.WaitForService("localhost:7233", time.Minute); err != nil {
		cancel()
		of.ErrorOutput(err.Error())
	}

	if err = handleTemporal(of); err != nil {
		cancel()
		of.ErrorOutput(err.Error())
	}

	for _, comp := range component.Components {
		if comp.Command == "bin/powerssl-temporalserver" {
			continue
		}
		if err = watcher.Add(comp.Command); err != nil {
			of.ErrorOutput(fmt.Sprintf("watcher error: %s", err))
		}
		addComponent(comp)
	}

	addComponent(component.Component{
		Name:    "grpcwebproxy",
		Command: "gobin",
		Args: strings.Join([]string{
			"-m",
			"-run",
			"github.com/improbable-eng/grpc-web/go/grpcwebproxy",
			"--allowed_origins http://localhost:8080",
			"--backend_addr localhost:8082",
			"--backend_tls",
			"--backend_tls_ca_files local/certs/ca.pem,local/certs/intermediate.pem",
			"--server_bind_address localhost",
			"--server_http_debug_port 8087",
			"--server_http_tls_port 8086",
			"--server_tls_cert_file local/certs/localhost.pem",
			"--server_tls_key_file local/certs/localhost-key.pem",
		}, " "),
	})

	addComponent(component.Component{
		Name:    "temporalweb",
		Command: "docker",
		Args: strings.Join([]string{
			"run", "--rm", "--init",
			"-e", "TEMPORAL_GRPC_ENDPOINT=host.docker.internal:7233",
			"-e", "TEMPORAL_TLS_CERT_PATH=/certs/localhost.pem",
			"-e", "TEMPORAL_TLS_KEY_PATH=/certs/localhost-key.pem",
			"-e", "TEMPORAL_TLS_CA_PATH=/certs/ca.pem",
			"-e", "TEMPORAL_TLS_SERVER_NAME=localhost",
			"-p", "8088:8088",
			"-v", fmt.Sprintf("%s/local/certs:/certs", wd),
			"temporalio/web:latest",
		}, " "),
	})

	if err = g.Wait(); err != nil {
		switch err.(type) {
		case util.InterruptError:
		default:
			of.ErrorOutput(fmt.Sprintf("error: %s", err))
		}
	}
}

func handlePostgres(of *internal.Outlet) error {
	{
		var err error
		var db *sql.DB
		if db, err = sql.Open("postgres", "postgresql://powerssl:powerssl@localhost:5432/?sslmode=disable"); err != nil {
			return errors.Wrap(err, "connecting default database")
		}
		defer func() {
			errWrapCloser(db, &err)
		}()
		for {
			if err = db.Ping(); err == nil {
				break
			}
			time.Sleep(time.Second)
		}
		of.SystemOutput("Create database: powerssl")
		if _, err = db.Exec("CREATE DATABASE powerssl;"); err != nil {
			if !strings.Contains(err.Error(), "already exists") {
				return errors.Wrap(err, "creating database powerssl")
			}
		}
		of.SystemOutput("Create database: temporal")
		if _, err = db.Exec("CREATE DATABASE temporal;"); err != nil {
			if !strings.Contains(err.Error(), "already exists") {
				return errors.Wrap(err, "creating database temporal")
			}
		}
		of.SystemOutput("Create database: vault")
		if _, err = db.Exec("CREATE DATABASE vault;"); err != nil {
			if !strings.Contains(err.Error(), "already exists") {
				return errors.Wrap(err, "creating database vault")
			}
		}
		errWrapCloser(db, &err)
	}
	{
		var err error
		var db *sql.DB
		if db, err = sql.Open("postgres", "postgresql://powerssl:powerssl@localhost:5432/vault?sslmode=disable"); err != nil {
			return errors.Wrap(err, "connecting vault database")
		}
		defer errWrapCloser(db, &err)
		for {
			if err = db.Ping(); err == nil {
				break
			}
			time.Sleep(time.Second)
		}
		of.SystemOutput("Create table: vault_kv_store")
		of.SystemOutput("Create index: parent_path_idx")
		if _, err = db.Exec("CREATE TABLE vault_kv_store(parent_path TEXT COLLATE \"C\" NOT NULL, path TEXT COLLATE \"C\", key TEXT COLLATE \"C\", value BYTEA, CONSTRAINT pkey PRIMARY KEY (path, key)); CREATE INDEX parent_path_idx ON vault_kv_store (parent_path);"); err != nil {
			if !strings.Contains(err.Error(), "already exists") {
				return errors.Wrap(err, "creating vault table and index")
			}
		}
		errWrapCloser(db, &err)
	}
	{
		comp := component.Component{
			Name:    "powerssl-apiserver",
			Command: "bin/powerutil",
			Args:    "migrate --database-url postgres://powerssl:powerssl@localhost:5432/powerssl?sslmode=disable up",
		}
		cmd, _, err := makeCmd(comp, 0, of)
		if err != nil {
			return err
		}
		if err = cmd.Start(); err != nil {
			return fmt.Errorf("failed to start %s: %s", comp.Command, err)
		}
		if err = cmd.Wait(); err != nil {
			return fmt.Errorf("failed to wait %s: %s", comp.Command, err)
		}
	}
	{
		comp := component.Component{
			Name:    "powerssl-temporalserver",
			Command: "bin/powerutil",
			Args:    "temporal migrate --host localhost --password powerssl --plugin postgres --port 5432 --user powerssl --docker",
		}
		cmd, _, err := makeCmd(comp, 0, of)
		if err != nil {
			return err
		}
		if err = cmd.Start(); err != nil {
			return fmt.Errorf("failed to start %s: %s", comp.Command, err)
		}
		if err = cmd.Wait(); err != nil {
			return fmt.Errorf("failed to wait %s: %s", comp.Command, err)
		}
	}
	return nil
}

func handleTemporal(of *internal.Outlet) error {
	comp := component.Component{
		Name:    "powerssl-temporalserver",
		Command: "bin/powerutil",
		Args:    "temporal register-namespace --address localhost:7233 --namespace powerssl --tls-cert-path local/certs/localhost.pem --tls-key-path local/certs/localhost-key.pem --tls-ca-path local/certs/ca.pem --tls-enable-host-verification --tls-server-name localhost",
	}
	cmd, _, err := makeCmd(comp, 0, of)
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return fmt.Errorf("failed to start %s: %s", comp.Command, err)
	}
	_ = cmd.Wait()
	return nil
}

func handleVault(of *internal.Outlet) error {
	var command, args string
	if _, err := os.Stat("local/vault/secret.yaml"); os.IsNotExist(err) {
		command = "bin/powerutil"
		args = "vault --ca local/certs/ca.pem --ca-key local/certs/ca-key.pem"
	} else {
		var byt []byte
		if byt, err = ioutil.ReadFile("local/vault/secret.yaml"); err != nil {
			return fmt.Errorf("config error: %s", err)
		}
		var config map[string]interface{}
		if err = yaml.Unmarshal(byt, &config); err != nil {
			return fmt.Errorf("config error: %s", err)
		}

		command = "vault"
		args = fmt.Sprintf("operator unseal -address https://localhost:8200 -ca-cert local/certs/ca.pem %s", config["keys"].([]interface{})[0].(string))
	}

	comp := component.Component{
		Name:    "vault",
		Command: command,
		Args:    args,
	}
	cmd, _, err := makeCmd(comp, 0, of)
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return fmt.Errorf("failed to start %s: %s", comp.Command, err)
	}
	if err = cmd.Wait(); err != nil {
		return fmt.Errorf("failed to wait %s: %s", comp.Command, err)
	}
	return nil
}

func serviceToken() (_ string, err error) {
	if err = internal.WaitForService("localhost:8081", time.Minute); err != nil {
		return "", err
	}
	var resp *http.Response
	if resp, err = http.Get("http://localhost:8081/service"); err != nil {
		return
	}
	defer errWrapCloser(resp.Body, &err)
	var byt []byte
	if byt, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return string(byt), nil
}

func errWrapCloser(closer io.Closer, wErr *error) {
	if err := closer.Close(); err != nil && *wErr != nil {
		*wErr = fmt.Errorf("%s: %w", err, *wErr)
	} else if err != nil {
		*wErr = err
	}
}
