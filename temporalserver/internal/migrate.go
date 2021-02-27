package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	_ "go.temporal.io/server/common/persistence/sql/sqlplugin/postgresql"
	"go.temporal.io/server/tools/sql"

	"powerssl.dev/common"
	"powerssl.dev/temporalserver/internal/migration"
)

const schemaVersion = "v96"

func RunMigrate(host, password, plugin, port, temporalDatabase, user, visibilityDatabase string) (err error) {
	commonArgs := []string{"--quiet", "--endpoint", host, "--port", port, "--user", user, "--password", password, "--plugin", plugin}

	var tempDir string
	if tempDir, err = ioutil.TempDir("", fmt.Sprintf("%s-", filepath.Base(os.Args[0]))); err != nil {
		return err
	}
	defer func() {
		if removeErr := os.RemoveAll(tempDir); removeErr != nil && err != nil {
			err = fmt.Errorf("%s: %w", removeErr, err)
		} else if removeErr != nil {
			err = removeErr
		}
	}()
	if err = migration.RestoreAssets(tempDir, ""); err != nil {
		return err
	}

	if err = runTemporalSQLTool(commonArgs, "create-database", "--database", temporalDatabase); err != nil && !strings.Contains(err.Error(), "already exists") {
		return err
	}
	if err = runTemporalSQLTool(commonArgs, "--database", temporalDatabase, "setup-schema", "--version", "0.0"); err != nil && !strings.Contains(err.Error(), "already exists") {
		return err
	}
	if err = runTemporalSQLTool(commonArgs, "--database", temporalDatabase, "update-schema", "--schema-dir", fmt.Sprintf("%s/schema/postgresql/%s/temporal/versioned", tempDir, schemaVersion)); err != nil {
		return err
	}

	if err = runTemporalSQLTool(commonArgs, "create-database", "--database", visibilityDatabase); err != nil && !strings.Contains(err.Error(), "already exists") {
		return err
	}
	if err = runTemporalSQLTool(commonArgs, "--database", visibilityDatabase, "setup-schema", "--version", "0.0"); err != nil && !strings.Contains(err.Error(), "already exists") {
		return err
	}
	if err = runTemporalSQLTool(commonArgs, "--database", visibilityDatabase, "update-schema", "--schema-dir", fmt.Sprintf("%s/schema/postgresql/%s/visibility/versioned", tempDir, schemaVersion)); err != nil {
		return err
	}

	return nil
}

func runTemporalSQLTool(commonArgs []string, args ...string) (err error) {
	var stderrFile, stdoutFile *os.File

	if stderrFile, err = ioutil.TempFile("", fmt.Sprintf("%s-", filepath.Base(os.Args[0]))); err != nil {
		return err
	}
	defer func() {
		if removeErr := os.Remove(stderrFile.Name()); removeErr != nil && err != nil {
			err = fmt.Errorf("%s: %w", removeErr, err)
		} else if removeErr != nil {
			err = removeErr
		}
	}()
	defer common.ErrWrapCloser(stderrFile, &err)

	if stdoutFile, err = ioutil.TempFile("", fmt.Sprintf("%s-", filepath.Base(os.Args[0]))); err != nil {
		return err
	}
	defer func() {
		if removeErr := os.Remove(stdoutFile.Name()); removeErr != nil && err != nil {
			err = fmt.Errorf("%s: %w", removeErr, err)
		} else if removeErr != nil {
			err = removeErr
		}
	}()
	defer common.ErrWrapCloser(stdoutFile, &err)

	stderr := os.Stderr
	os.Stderr = stderrFile
	stdout := os.Stdout
	os.Stdout = stdoutFile

	err = sql.RunTool(append([]string{"temporal-sql-tool"}, append(commonArgs, args...)...))

	os.Stderr = stderr
	os.Stdout = stdout

	var stderrContent, stdoutContent []byte
	if stderrContent, err = ioutil.ReadFile(stderrFile.Name()); err != nil {
		return err
	}
	if stdoutContent, err = ioutil.ReadFile(stdoutFile.Name()); err != nil {
		return err
	}

	if len(stderrContent) > 0 {
		return fmt.Errorf(string(stderrContent))
	}

	fmt.Println(string(stdoutContent))

	return err
}
