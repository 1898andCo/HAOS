package module

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/sirupsen/logrus"
	"pault.ag/go/modprobe"
)

const (
	procModulesFile = "/proc/modules"
)

type Abstract interface {
	Open(string) (*os.File, error)
	Close(*os.File) error
	Load(string, string) error
}

type Concrete struct{}

func (Concrete) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (Concrete) Close(f *os.File) error {
	return f.Close()
}
func (Concrete) Load(module, params string) error {
	return modprobe.Load(module, params)
}
func LoadModules(cfg *config.CloudConfig, m Abstract) error {
	loaded := map[string]bool{}
	f, err := m.Open(procModulesFile)
	if err != nil {
		return err
	}
	defer m.Close(f)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		loaded[strings.SplitN(sc.Text(), " ", 2)[0]] = true
	}
	modules := cfg.HAOS.Modules
	for _, m := range modules {
		if loaded[m] {
			continue
		}
		params := strings.SplitN(m, " ", -1)
		logrus.Debugf("module %s with parameters [%s] is loading", m, params)
		if err := modprobe.Load(params[0], strings.Join(params[1:], " ")); err != nil {
			return fmt.Errorf("could not load module %s with parameters [%s], err %v", m, params, err)
		}
		logrus.Debugf("module %s is loaded", m)
	}
	return sc.Err()
}
