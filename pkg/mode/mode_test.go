package mode_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/mocks"
	"github.com/1898andCo/HAOS/pkg/mode"
	"github.com/spf13/afero"
)

func TestGet(t *testing.T) {
	cfg := mocks.NewCloudConfig()
	afero.WriteFile(cfg.Fs, "/run/HAOS/mode", []byte("install"), 0644)
	s, err := mode.Get(cfg)
	if err != nil {
		t.Errorf("Get() error = %v", err)
	}
	if s != "install" {
		t.Errorf("Get() s = %v", s)
	}
	s, err = mode.Get(cfg, "/path/does/not/exist")
	if s != "" {
		t.Errorf("Get() s = %v", s)
	}
	if err != nil {
		t.Errorf("Get() error = %v", err)
	}
}
