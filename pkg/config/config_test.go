package config_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/config"
)

func TestPermissions(t *testing.T) {
	f := config.File{
		RawFilePermissions: "0644",
	}
	perms, err := f.Permissions()
	if err != nil {
		t.Errorf("Permissions() error = %v", err)
	}
	if perms != 0644 {
		t.Errorf("Permissions() perms = %v", perms)
	}
}
