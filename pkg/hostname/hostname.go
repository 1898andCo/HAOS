// Package hostname is responsible for setting the hostname of the system and syncing /etc/hosts
//
// The hostname is configured in the passed in configuration struct.

package hostname

import (
	"bufio"
	"os"
	"strings"
	"syscall"

	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/1898andCo/HAOS/pkg/system"
	
	"github.com/spf13/afero"
)

func SetHostname(c *config.CloudConfig) error {
	hostname := c.Hostname
	if hostname == "" {
		return nil
	}
	if err := syscall.Sethostname([]byte(hostname)); err != nil {
		return err
	}
	return syncHostname()
}

func syncHostname() error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	if hostname == "" {
		return nil
	}

	if err := afero.WriteFile(system.AppFs, "/etc/hostname", []byte(hostname+"\n"), 0644); err != nil {
		return err
	}

	hosts, err := os.Open("/etc/hosts")
	defer hosts.Close()
	if err != nil {
		return err
	}
	lines := bufio.NewScanner(hosts)
	content := ""
	for lines.Scan() {
		line := strings.TrimSpace(lines.Text())
		fields := strings.Fields(line)
		if len(fields) > 0 && fields[0] == "127.0.1.1" {
			content += "127.0.1.1 " + hostname + "\n"
			continue
		}
		content += line + "\n"
	}
	return afero.WriteFile(system.AppFs, "/etc/hosts", []byte(content), 0600)
}
