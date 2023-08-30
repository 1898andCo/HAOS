package writefile_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/1898andCo/HAOS/pkg/mocks"
	"github.com/1898andCo/HAOS/pkg/system"
	"github.com/1898andCo/HAOS/pkg/writefile"
	"github.com/spf13/afero"
)

func setupFS() {
	system.AppFs = afero.NewMemMapFs()
}

func TestWriteFiles(t *testing.T) {
	setupFS()
	cfg := mocks.NewCloudConfig()
	cfg.WriteFiles = []config.File{
		{
			Path:               "/etc/test",
			Content:            "test",
			Owner:              "root",
			RawFilePermissions: "0644",
		},
	}
	writefile.WriteFiles(cfg)
}
