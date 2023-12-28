package hostname

import (
	"os"
	"testing"

	"github.com/1898andCo/HAOS/pkg/config"
)

type HostnameMock struct {
	SethostnameFunc func([]byte) error
}

func (HostnameMock) Sethostname(hostname []byte) error {
	return nil
}

func (HostnameMock) Hostname() (string, error) {
	return "test", nil
}

func (HostnameMock) Close(file *os.File) error {
	return file.Close()
}

func (HostnameMock) WriteFile(path string, data []byte, perm os.FileMode) error {
	return nil
}

const hostsFilecontent (string) = `127.0.1.1
127.0.0.1
0.0.0.0`

const hostnameFilecontent (string) = `test`

func (HostnameMock) Open(path string) (*os.File, error) {
	file := os.NewFile(0, "test")
	file.WriteString(hostnameFilecontent)
	return file, nil
}

func TestSetHostname(t *testing.T) {
	c := &config.CloudConfig{
		Hostname: "test",
	}
	call := &HostnameMock{}
	err := SetHostname(c, call)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
