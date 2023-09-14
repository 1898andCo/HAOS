package cliinstall_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/cliinstall"
	"github.com/1898andCo/HAOS/pkg/mocks"
	"github.com/1898andCo/HAOS/pkg/system"
	"github.com/spf13/afero"
)

func TestAsk(t *testing.T) {
	system.AppFs = afero.NewMemMapFs()
	afero.WriteFile(system.AppFs, "mode", []byte("install"), 0644)
	ask, err := cliinstall.Ask(mocks.NewCloudConfig())
	if err != nil {
		t.Errorf("Ask() error = %v", err)
	}
	if ask != false {
		t.Errorf("Ask() ask = %v", ask)
	}
}
