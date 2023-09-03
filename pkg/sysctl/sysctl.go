// Package sysctl is responsible for configuring sysctl settings
package sysctl

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/1898andCo/HAOS/pkg/config"
)

func ConfigureSysctl(cfg *config.CloudConfig) error {
	for k, v := range cfg.HAOS.Sysctls {
		elements := []string{"/proc", "sys"}
		elements = append(elements, strings.Split(k, ".")...)
		path := path.Join(elements...)
		if err := ioutil.WriteFile(path, []byte(v), 0644); err != nil {
			return err
		}
	}
	return nil
}
