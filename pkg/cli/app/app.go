package app

import (
	"fmt"

	"github.com/BOHICA-LABS/BLAOS/pkg/cli/config"
	"github.com/BOHICA-LABS/BLAOS/pkg/cli/install"
	"github.com/BOHICA-LABS/BLAOS/pkg/cli/rc"
	"github.com/BOHICA-LABS/BLAOS/pkg/cli/upgrade"
	"github.com/BOHICA-LABS/BLAOS/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	Debug bool
)

// New CLI App
func New() *cli.App {
	app := cli.NewApp()
	app.Name = "BLAOS"
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
			EnvVar:      "BLAOS_DEBUG",
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
