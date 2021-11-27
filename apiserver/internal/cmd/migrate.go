package cmd

import (
	"errors"
	"fmt"
	"powerssl.dev/apiserver/internal/db"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spangenberg/snakecharmer"
	"github.com/spf13/cobra"
)

func newCmdMigrate() *cobra.Command {
	var databaseURL string
	var m *migrate.Migrate

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if databaseURL == "" {
				return errors.New("database URL can't be blank")
			}
			var err error
			var sourceInstance source.Driver
			if sourceInstance, err = iofs.New(db.MigrationsFS, "migrations"); err != nil {
				return err
			}
			m, err = migrate.NewWithSourceInstance("go-bindata", sourceInstance, databaseURL)
			return err
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			if m == nil {
				return nil
			}
			sourceErr, databaseErr := m.Close()
			var errStrings []string
			if sourceErr != nil {
				errStrings = append(errStrings, sourceErr.Error())
			}
			if databaseErr != nil {
				errStrings = append(errStrings, databaseErr.Error())
			}
			if len(errStrings) > 0 {
				return errors.New(strings.Join(errStrings, "\n"))
			}
			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&databaseURL, "database-url", "", "Database URL")

	cmd.AddCommand(newCmdMigrateDown(&m))
	cmd.AddCommand(newCmdMigrateDrop(&m))
	cmd.AddCommand(newCmdMigrateForce(&m))
	cmd.AddCommand(newCmdMigrateGoto(&m))
	cmd.AddCommand(newCmdMigrateUp(&m))

	return cmd
}

func newCmdMigrateDown(m **migrate.Migrate) *cobra.Command {
	var all bool
	var limit = -1

	cmd := &cobra.Command{
		Use:   "down [--all | -a | N]",
		Short: "Apply all or N down migrations",
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if all && len(args) > 0 {
				return errors.New("--all cannot be used with other arguments")
			}
			if !all && len(args) == 0 {
				cmd.Println("Are you sure you want to apply all down migrations? [y/N]")
				var response string
				if _, err := fmt.Scanln(&response); err != nil {
					return err
				}
				response = strings.ToLower(strings.TrimSpace(response))
				if response == "y" {
					cmd.Println("Applying all down migrations")
				} else {
					cmd.Println("Not applying all down migrations")
				}
				return nil
			}
			if !all {
				n, err := strconv.ParseUint(args[0], 10, 64)
				if err != nil {
					return errors.New("can't read limit argument N")
				}
				limit = int(n)
				return nil
			}
			return nil
		},
		Run: snakecharmer.HandleError(func(cmd *cobra.Command, args []string) error {
			if limit >= 0 {
				if err := (*m).Steps(-limit); err != nil {
					if err != migrate.ErrNoChange {
						return err
					}
				}
			} else {
				if err := (*m).Down(); err != nil {
					if err != migrate.ErrNoChange {
						return err
					}
				}
			}
			return nil
		}),
	}

	cmd.Flags().BoolVarP(&all, "all", "a", false, "Apply all down migrations")

	return cmd
}

func newCmdMigrateDrop(m **migrate.Migrate) *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "drop [--force | -f]",
		Short: "Drop everything inside database",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !force {
				cmd.Println("Are you sure you want to drop the entire database schema? [y/N]")
				var response string
				if _, err := fmt.Scanln(&response); err != nil {
					return err
				}
				response = strings.ToLower(strings.TrimSpace(response))
				if response == "y" {
					cmd.Println("Dropping the entire database schema")
				} else {
					cmd.Println("Aborted dropping the entire database schema")
					return errors.New("aborted")
				}
			}
			return nil
		},
		Run: snakecharmer.HandleError(func(cmd *cobra.Command, args []string) error {
			return (*m).Drop()
		}),
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Use to bypass confirmation")

	return cmd
}

func newCmdMigrateForce(m **migrate.Migrate) *cobra.Command {
	var v int

	cmd := &cobra.Command{
		Use:   "force V",
		Short: "Set version V but don't run migration (ignores dirty state)",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("please specify version argument V")
			}
			v64, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return errors.New("can't read version argument V")
			}
			if v64 < -1 {
				return errors.New("argument V must be >= -1")
			}
			v = int(v64)
			return nil
		},
		Run: snakecharmer.HandleError(func(cmd *cobra.Command, args []string) error {
			return (*m).Force(v)
		}),
	}

	return cmd
}

func newCmdMigrateGoto(m **migrate.Migrate) *cobra.Command {
	var v uint

	cmd := &cobra.Command{
		Use:   "goto V",
		Short: "Migrate to version V",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("please specify version argument V")
			}
			v64, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return errors.New("can't read version argument V")
			}
			v = uint(v64)
			return nil
		},
		Run: snakecharmer.HandleError(func(cmd *cobra.Command, args []string) error {
			if err := (*m).Migrate(v); err != nil {
				if err != migrate.ErrNoChange {
					return err
				}
			}
			return nil
		}),
	}

	return cmd
}

func newCmdMigrateUp(m **migrate.Migrate) *cobra.Command {
	var limit = -1

	cmd := &cobra.Command{
		Use:   "up [N]",
		Short: "Apply all or N up migrations",
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				var n uint64
				var err error
				if n, err = strconv.ParseUint(args[0], 10, 64); err != nil {
					return errors.New("can't read limit argument N")
				}
				limit = int(n)
			}
			return nil
		},
		Run: snakecharmer.HandleError(func(cmd *cobra.Command, args []string) error {
			if limit >= 0 {
				if err := (*m).Steps(limit); err != nil {
					if err != migrate.ErrNoChange {
						return err
					}
				}
			} else {
				if err := (*m).Up(); err != nil {
					if err != migrate.ErrNoChange {
						return err
					}
				}
			}
			return nil
		}),
	}

	return cmd
}
