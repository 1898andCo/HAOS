package module

import (
	"os"
	"testing"

	"github.com/1898andCo/HAOS/pkg/config"
)

type Mock struct{}

func (Mock) Open(name string) (*os.File, error) {
	return os.NewFile(0, "mock"), nil
}

func (Mock) Load(module, params string) error {
	return nil
}

func (Mock) Close(f *os.File) error {
	err := f.Close()
	if err != nil {
		return err
	}
	return os.Remove(f.Name())
}

func TestModule(t *testing.T) {
	m := Mock{}
	err := LoadModules(&config.CloudConfig{HAOS: config.HAOS{Modules: []string{}}}, m)
	if err != nil {
		t.Errorf("failed to load modules: %v", err)
	}
}
