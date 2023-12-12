// Package install is responsible for running the install phase of the system
//
// it exposes a single function Command() which returns a configured cli.Command struct
package install

import (
	"fmt"
	"os"

	"github.com/1898andCo/HAOS/pkg/cliinstall"
	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/1898andCo/HAOS/pkg/mode"
	"github.com/1898andCo/HAOS/pkg/system"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func Command() cli.Command {
	cfg := &config.CloudConfig{
		Fs: system.AppFs,
	}
	mode, _ := mode.Get(cfg)
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
