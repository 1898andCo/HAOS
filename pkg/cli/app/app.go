// Package app
//
// It is the entrypoint for the HAOS command line interface. It is responsible for
// initializing the CLI app and its subcommands. It also handles the global flags passed
// to the CLI app.
package app

import (
	"fmt"

	"github.com/1898andCo/HAOS/pkg/cli/config"
	"github.com/1898andCo/HAOS/pkg/cli/install"
	"github.com/1898andCo/HAOS/pkg/cli/rc"
	"github.com/1898andCo/HAOS/pkg/cli/upgrade"
	"github.com/1898andCo/HAOS/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	Debug bool
)

// New CLI App
func New() *cli.App {
	app := cli.NewApp()
	app.Name = "haos"
	app.Usage = "Booting to k3s so you don't have to"
	app.Version = version.Version
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version %s\n", app.Name, app.Version)
	}
	// required flags without defaults will break symlinking to exe with name of sub-command as target
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Turn on debug logs",
			EnvVar:      "HAOS_DEBUG",
			Destination: &Debug,
		},
	}

	app.Commands = []cli.Command{
		rc.Command(),
		config.Command(),
		install.Command(),
		upgrade.Command(),
	}

	app.Before = func(c *cli.Context) error {
		if Debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	return app
}
