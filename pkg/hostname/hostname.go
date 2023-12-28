package hostname

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"github.com/1898andCo/HAOS/pkg/config"
)

type HostnameAbstract interface {
	Sethostname([]byte) error
	Hostname() (string, error)
	Open(string) (*os.File, error)
	Close(*os.File) error
	WriteFile(string, []byte, os.FileMode) error
}

type HostnameConcrete struct{}

func (HostnameConcrete) Sethostname(hostname []byte) error {
	return syscall.Sethostname(hostname)
}

func (HostnameConcrete) Hostname() (string, error) {
	return os.Hostname()
}

func (HostnameConcrete) Open(path string) (*os.File, error) {
	return os.Open(path)
}

func (HostnameConcrete) Close(file *os.File) error {
	return file.Close()
}

func (HostnameConcrete) WriteFile(path string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(path, data, perm)
}

func SetHostname(c *config.CloudConfig, call HostnameAbstract) error {
	hostname := c.Hostname
	if hostname == "" {
		return nil
	}
	if err := call.Sethostname([]byte(hostname)); err != nil {
		return err
	}
	return syncHostname(call)
}

func syncHostname(h HostnameAbstract) error {
	hostname, err := h.Hostname()
	if err != nil {
		return err
	}
	if hostname == "" {
		return nil
	}

	if err := h.WriteFile("/etc/hostname", []byte(hostname+"\n"), 0644); err != nil {
		return err
	}

	hosts, err := h.Open("/etc/hosts")
	defer h.Close(hosts)
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
	return h.WriteFile("/etc/hosts", []byte(content), 0600)
}
