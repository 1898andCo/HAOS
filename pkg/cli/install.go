// Package install is responsible for running the install phase of the system
//
// it exposes a single function Command() which returns a configured cli.Command struct
package cli

import (
	"fmt"
	"os"

	"github.com/1898andCo/HAOS/pkg/cliinstall"
	"github.com/1898andCo/HAOS/pkg/mode"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func InstallCommand() cli.Command {
	mode, _ := mode.Get()
	return cli.Command{
		Name:  "install",
		Usage: "install HAOS",
		Flags: []cli.Flag{},
		Before: func(c *cli.Context) error {
			if os.Getuid() != 0 {
				return fmt.Errorf("must be run as root")
			}
			return nil
		},
		Action: func(*cli.Context) {
			if err := cliinstall.Run(); err != nil {
				logrus.Error(err)
			}
		},
		Hidden: mode == "local",
	}
}