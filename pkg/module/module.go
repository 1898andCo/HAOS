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

type abstract interface {
	Open(string) (*os.File, error)
	Close(*os.File) error
	Load(string, string) error
}

type concrete struct{}

func (concrete) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (concrete) Close(f *os.File) error {
	return f.Close()
}

func (concrete) Load(module, params string) error {
	return modprobe.Load(module, params)
}

var impl abstract = concrete{}

func LoadModules(cfg *config.CloudConfig) error {
	loaded := map[string]bool{}
	f, err := impl.Open(procModulesFile)
	if err != nil {
		return err
	}
	defer impl.Close(f)
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
