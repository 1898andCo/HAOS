package mode_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/mode"
	"github.com/1898andCo/HAOS/pkg/system"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

func TestGet(t *testing.T) {
	system.AppFs = &afero.MemMapFs{}
	afero.WriteFile(system.AppFs, "/run/HAOS/mode", []byte("install"), 0644)
	s, err := mode.Get()
	logrus.Info(s)
	if err != nil {
		t.Errorf("Get() error = %v", err)
	}
	if s != "install" {
		t.Errorf("Get() s = %v", s)
	}
	s, err = mode.Get("/path/does/not/exist")
	if s != "" {
		t.Errorf("Get() s = %v", s)
	}
	if err != nil {
		t.Errorf("Get() error = %v", err)
	}
}
