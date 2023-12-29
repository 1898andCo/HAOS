package hostname

import (
	"os"
	"testing"

	"github.com/1898andCo/HAOS/pkg/config"
)

type Mock struct {
	SethostnameFunc func([]byte) error
}

func (Mock) Sethostname(hostname []byte) error {
	return nil
}

func (Mock) Hostname() (string, error) {
	return "test", nil
}

func (Mock) Close(file *os.File) error {
	return file.Close()
}

func (Mock) WriteFile(path string, data []byte, perm os.FileMode) error {
	return nil
}

const hostsFilecontent (string) = `127.0.1.1
127.0.0.1
0.0.0.0`

const hostnameFilecontent (string) = `test`

func (Mock) Open(path string) (*os.File, error) {
	file := os.NewFile(0, "test")
	file.WriteString(hostsFilecontent)
	return file, nil
}

func TestSetHostname(t *testing.T) {
	c := &config.CloudConfig{
		Hostname: "test",
	}
	call := &Mock{}
	err := SetHostname(c, call)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
