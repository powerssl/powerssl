package internal

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"powerssl.dev/common/errutil"
	"powerssl.dev/common/interrupthandler"
	"powerssl.dev/common/log"

	"powerssl.dev/tools/dev-runner/internal/component"
)

func Run() error {
	var g *errgroup.Group
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx = errgroup.WithContext(ctx)
	g.Go(func() error {
		var logger log.Logger
		var err error
		if logger, err = log.New(log.Config{
			Env: "production",
		}); err != nil {
			return err
		}
		defer errutil.ErrWrapSync(logger, &err)
		return interrupthandler.InterruptHandler(ctx, logger)
	})

	go func() {
		_ = http.ListenAndServe("localhost:8080", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			target := "https://" + strings.ReplaceAll(req.Host, "8080", "8443") + req.URL.Path
			if len(req.URL.RawQuery) > 0 {
				target += "?" + req.URL.RawQuery
			}
			http.Redirect(w, req, target, http.StatusTemporaryRedirect)
		}))
	}()

	of := NewOutlet()
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
		defer errutil.ErrWrapCloser(watcher, &err)
	}

	interrupts := make(map[string]chan struct{}, len(component.Components)+1)

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create != fsnotify.Create {
					continue
				}
				if c, ok := interrupts[event.Name]; ok {
					c <- struct{}{}
				}
			case err := <-watcher.Errors:
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
				for k, v := range localComp.Env {
					switch v {
					case component.GenerateAuthToken:
						var err error
						if localComp.Env[k], err = serviceToken(comp.Name); err != nil {
							of.SystemOutput(err.Error())
							cancel()
						}
					case component.GenerateVaultAppRoleID:
						var err error
						if localComp.Env[k], err = vaultRoleID(comp.Name); err != nil {
							of.SystemOutput(err.Error())
							cancel()
						}
					case component.GenerateVaultAppSecretID:
						var err error
						if localComp.Env[k], err = vaultSecretID(comp.Name); err != nil {
							of.SystemOutput(err.Error())
							cancel()
						}
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
		Args: []string{
			"run", "--rm",
			"-e", "POSTGRES_PASSWORD=powerssl",
			"-e", "POSTGRES_DB=powerssl",
			"-e", "POSTGRES_USER=powerssl",
			"-e", "PGDATA=/var/lib/postgresql/data/pgdata",
			"-p", "5432:5432",
			"-v", wd + "/local/postgresql/data:/var/lib/postgresql/data",
			"postgres:13.1",
		},
	})

	if err = waitForService("localhost:5432", time.Minute); err != nil {
		of.SystemOutput(err.Error())
		cancel()
	}

	if err = handlePostgres(of); err != nil {
		cancel()
		of.ErrorOutput(err.Error())
	}

	addComponent(component.Component{
		Name:    "vault",
		Command: "docker",
		Args: []string{
			"run", "--rm", "--cap-add=IPC_LOCK",
			"-v", wd + "/tools/dev-runner/configs:/vault/config",
			"-v", wd + "/local/certs/localhost.pem:/etc/ssl/certs/localhost.pem",
			"-v", wd + "/local/certs/localhost-key.pem:/etc/ssl/private/localhost-key.pem",
			"-p", "8200:8200",
			"vault:1.7.0",
			"server",
		},
	})

	if err = waitForService("localhost:8200", time.Minute); err != nil {
		cancel()
		of.ErrorOutput(err.Error())
	}

	if err = handleVault(of); err != nil {
		cancel()
		of.ErrorOutput(err.Error())
	}

	for _, comp := range component.Components {
		if comp.Name != "powerssl-temporal" {
			continue
		}
		if err = watcher.Add(comp.Command); err != nil {
			of.ErrorOutput(fmt.Sprintf("watcher error: %s", err))
		}
		addComponent(comp)
	}

	if err = waitForService("localhost:7233", time.Minute); err != nil {
		cancel()
		of.ErrorOutput(err.Error())
	}

	if err = handleTemporal(of); err != nil {
		cancel()
		of.ErrorOutput(err.Error())
	}

	for _, comp := range component.Components {
		if comp.Name == "powerssl-temporal" {
			continue
		}
		if err = watcher.Add(comp.Command); err != nil {
			of.ErrorOutput(fmt.Sprintf("watcher error: %s", err))
		}
		addComponent(comp)
	}

	addComponent(component.Component{
		Name:    "grpcwebproxy",
		Command: "go",
		Args: []string{
			"run",
			"github.com/improbable-eng/grpc-web/go/grpcwebproxy@latest",
			"--allowed_origins", "https://localhost:8443",
			"--backend_addr", "localhost:8082",
			"--backend_tls",
			"--backend_tls_ca_files", "local/certs/ca.pem,local/certs/intermediate.pem",
			"--server_bind_address", "localhost",
			"--server_http_debug_port", "8889",
			"--server_http_tls_port", "8883",
			"--server_tls_cert_file", "local/certs/localhost.pem",
			"--server_tls_key_file", "local/certs/localhost-key.pem",
		},
	})

	addComponent(component.Component{
		Name:    "temporalweb",
		Command: "docker",
		Args: []string{
			"run", "--rm", "--init",
			"-e", "TEMPORAL_GRPC_ENDPOINT=host.docker.internal:7233",
			"-e", "TEMPORAL_TLS_CERT_PATH=/certs/localhost.pem",
			"-e", "TEMPORAL_TLS_KEY_PATH=/certs/localhost-key.pem",
			"-e", "TEMPORAL_TLS_CA_PATH=/certs/ca.pem",
			"-e", "TEMPORAL_TLS_SERVER_NAME=localhost",
			"-p", "8088:8088",
			"-v", wd + "/local/certs:/certs",
			"temporalio/web:latest",
		},
	})

	if err = g.Wait(); err != nil {
		switch err.(type) {
		case interrupthandler.InterruptError:
		default:
			return err
		}
	}
	return nil
}

func handlePostgres(of *Outlet) error {
	{
		var err error
		var db *sql.DB
		if db, err = sql.Open("postgres", "postgresql://powerssl:powerssl@localhost:5432/?sslmode=disable"); err != nil {
			return errors.Wrap(err, "connecting default database")
		}
		defer func() {
			errutil.ErrWrapCloser(db, &err)
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
		errutil.ErrWrapCloser(db, &err)
	}
	{
		var err error
		var db *sql.DB
		if db, err = sql.Open("postgres", "postgresql://powerssl:powerssl@localhost:5432/vault?sslmode=disable"); err != nil {
			return errors.Wrap(err, "connecting vault database")
		}
		defer errutil.ErrWrapCloser(db, &err)
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
		errutil.ErrWrapCloser(db, &err)
	}
	{
		comp := component.Component{
			Name:    "powerssl-apiserver",
			Command: "bin/powerssl-apiserver",
			Args: []string{
				"migrate",
				"--database-url", "postgres://powerssl:powerssl@localhost:5432/powerssl?sslmode=disable",
				"up",
			},
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
			Name:    "powerssl-temporal",
			Command: "bin/powerssl-temporal",
			Args: []string{
				"migrate",
				"--host", "localhost",
				"--password", "powerssl",
				"--plugin", "postgres",
				"--port", "5432",
				"--user", "powerssl",
			},
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

func handleTemporal(of *Outlet) error {
	comp := component.Component{
		Name:    "powerssl-temporal",
		Command: "bin/powerssl-temporal",
		Args: []string{
			"register-namespace",
			"--address", "localhost:7233",
			"--namespace", "powerssl",
			"--tls-cert-path", "local/certs/localhost.pem",
			"--tls-key-path", "local/certs/localhost-key.pem",
			"--tls-ca-path", "local/certs/ca.pem",
			"--tls-enable-host-verification",
			"--tls-server-name", "localhost",
		},
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

func serviceToken(_ string) (_ string, err error) {
	if err = waitForService("localhost:8843", time.Minute); err != nil {
		return "", err
	}
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}
	localCertFile := "local/certs/ca.pem" // TODO: make const
	certs, err := ioutil.ReadFile(localCertFile)
	if err != nil {
		return "", fmt.Errorf("failed to append %q to RootCAs: %v", localCertFile, err)
	}
	rootCAs.AppendCertsFromPEM(certs)
	tlsConfig := &tls.Config{
		ServerName: "localhost",
		RootCAs:    rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: tlsConfig}
	httpClient := &http.Client{Transport: tr}
	var resp *http.Response
	if resp, err = httpClient.Get("https://localhost:8843/service"); err != nil {
		return
	}
	defer errutil.ErrWrapCloser(resp.Body, &err)
	var byt []byte
	if byt, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	return string(byt), nil
}
