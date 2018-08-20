package main

import (
	"os"
	"os/user"

	"gopkg.in/urfave/cli.v2"
	// "powerssl.io/pkg/ctl/cmd"
)

/*
var (
	config  = kingpin.Flag("config", "Config file path.").Short('c').Default(defaultConfig()).ExistingFile()
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()

	describeDomains   = kingpin.Command("describe-domains", "Describe domains.")
	describeDomainsId = describeDomains.Flag("id", "IDs to describe.").String()

	describeCertificates   = kingpin.Command("describe-certificates", "Describe certificates.")
	describeCertificatesId = describeCertificates.Flag("id", "IDs to describe.").String()
)
*/

func main() {
	app := &cli.App{
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			{
				Name:     "create-domain",
				Category: "Domain service",
			},
			{
				Name:     "delete-domain",
				Category: "Domain service",
			},
			{
				Name:     "get-domain",
				Category: "Domain service",
			},
			{
				Name:     "list-domains",
				Category: "Domain service",
				Action: func(c *cli.Context) error {
					// commands.ListDomains()
					return nil
				},
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "show-discarded"},
					&cli.StringFlag{Name: "next-page-token"},
					&cli.UintFlag{Name: "page-size"},
				},
			},
			{
				Name:     "update-domain",
				Category: "Domain service",
			},
		},
		Name:    "powersslctl",
		Version: "0.1.0",
	}

	app.Run(os.Args)
}

func defaultConfig() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	return user.HomeDir + "/.powerssl/config.json"
}
