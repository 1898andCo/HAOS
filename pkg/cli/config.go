// Package config
//

package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/1898andCo/HAOS/pkg/cc"
	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	initrd       = false
	bootPhase    = false
	installPhase = false
	dump         = false
	dumpJSON     = false
)

// Command returns a configured cli.Command struct
func ConfigCommand() cli.Command {
	return cli.Command{
		Name:      "config",
		Usage:     "configure HAOS",
		ShortName: "cfg",
		// Aliases: []string{
		// 	"ccapply",
		// },
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:        "initrd",
				Destination: &initrd,
				Usage:       "Run initrd stage",
			},
			cli.BoolFlag{
				Name:        "boot",
				Destination: &bootPhase,
				Usage:       "Run boot stage",
			},
			cli.BoolFlag{
				Name:        "install",
				Destination: &installPhase,
				Usage:       "Run install stage",
			},
			cli.BoolFlag{
				Name:        "dump",
				Destination: &dump,
				Usage:       "Print current configuration",
			},
			cli.BoolFlag{
				Name:        "dump-json",
				Destination: &dumpJSON,
				Usage:       "Print current configuration in json",
			},
		},
		Before: func(c *cli.Context) error {
			if os.Getuid() != 0 {
				return fmt.Errorf("must be run as root")
			}
			return nil
		},
		Action: func(*cli.Context) {
			if err := Main(); err != nil {
				logrus.Error(err)
			}
		},
	}
}

// Main is the main function for the config command
func Main() error {
	cfg, err := config.ReadConfig()
	if err != nil {
		return err
	}

	if initrd {
		return cc.InitApply(&cfg)
	} else if bootPhase {
		return cc.BootApply(&cfg)
	} else if installPhase {
		return cc.InstallApply(&cfg)
	} else if dump {
		return config.Write(cfg, os.Stdout)
	} else if dumpJSON {
		return json.NewEncoder(os.Stdout).Encode(&cfg)
	}

	return cc.RunApply(&cfg)
}