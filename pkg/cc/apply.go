// Package cc is responsible for applying the cloud config to the system
//
// see [pkg/config](../config) for more information on the cloud config
// TODO: improve documentation
package cc

import (
	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/urfave/cli"
)

// applier abstracts a function call that applies a cloud config
// to a given config struct
type applier func(cfg *config.CloudConfig) error

// runApplies runs a list of appliers and returns an error if any of them fail
func runApplies(cfg *config.CloudConfig, appliers ...applier) error {
	var errors []error

	for _, a := range appliers {
		err := a(cfg)
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return cli.NewMultiError(errors...)
	}

	return nil
}

// RunApply runs all current configuration 'run' functions using the passed in config
func RunApply(cfg *config.CloudConfig) error {
	return runApplies(cfg,
		ApplyModules,
		ApplySysctls,
		ApplyHostname,
		ApplyDNS,
		ApplyWifi,
		ApplyPassword,
		ApplySSHKeysWithNet,
		ApplyWriteFiles,
		ApplyBootManifests,
		ApplyEnvironment,
		ApplyRuncmd,
		ApplyInstall,
		ApplyK3SInstall,
	)
}

// InstallApply runs all current configuration 'install' functions using the passed in config
func InstallApply(cfg *config.CloudConfig) error {
	return runApplies(cfg,
		ApplyK3SWithRestart,
	)
}

// BootApply runs all current configuration 'boot' functions using the passed in config
func BootApply(cfg *config.CloudConfig) error {
	return runApplies(cfg,
		ApplyDataSource,
		ApplyModules,
		ApplySysctls,
		ApplyHostname,
		ApplyDNS,
		ApplyWifi,
		ApplyPassword,
		ApplySSHKeys,
		ApplyK3SNoRestart,
		ApplyWriteFiles,
		ApplyEnvironment,
		ApplyBootcmd,
	)
}

// InitApply runs all current configuration 'init' functions using the passed in config
func InitApply(cfg *config.CloudConfig) error {
	return runApplies(cfg,
		ApplyModules,
		ApplySysctls,
		ApplyHostname,
		ApplyWriteFiles,
		ApplyEnvironment,
		ApplyInitcmd,
	)
}
