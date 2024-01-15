package module

import (
	"os"
	"testing"

	"github.com/1898andCo/HAOS/pkg/config"
)

type mock struct{}

func (mock) Open(name string) (*os.File, error) {
	return os.NewFile(0, "mock"), nil
}

func (mock) Load(module, params string) error {
	return nil
}

func (mock) Close(f *os.File) error {
	err := f.Close()
	if err != nil {
		return err
	}
	return os.Remove(f.Name())
}

func TestModule(t *testing.T) {
	impl = mock{}
	err := LoadModules(&config.CloudConfig{HAOS: config.HAOS{Modules: []string{}}})
	if err != nil {
		t.Errorf("failed to load modules: %v", err)
	}
}
