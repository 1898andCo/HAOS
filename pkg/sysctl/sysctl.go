// Package sysctl is responsible for configuring sysctl settings
package sysctl

import (
	"path"
	"strings"

	"github.com/spf13/afero"
	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/1898andCo/HAOS/pkg/system"
)

func ConfigureSysctl(cfg *config.CloudConfig) error {
	for k, v := range cfg.HAOS.Sysctls {
		elements := []string{"/proc", "sys"}
		elements = append(elements, strings.Split(k, ".")...)
		path := path.Join(elements...)
		if err := afero.WriteFile(system.AppFs, path, []byte(v), 0644); err != nil {
			return err
		}
	}
	return nil
}
