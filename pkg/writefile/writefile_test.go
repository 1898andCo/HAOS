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
	system.AppFs.MkdirAll("/tmp", 0755)
	cfg := mocks.NewCloudConfig()
	cfg.WriteFiles = []config.File{
		{
			Path:    "/tmp/test",
			Content: "test",
		},
	}
	writefile.WriteFiles(cfg)
	t.Logf("cfg: %v", cfg.WriteFiles)
	// test to see if the files got written to the in-mem filesystem
	content, err := afero.ReadFile(system.AppFs, cfg.WriteFiles[0].Path)
	t.Logf("content: %s", content)
	if err != nil {
		t.Errorf("failed to read file: %v", err)
	}
	if string(content) != cfg.WriteFiles[0].Content {
		t.Errorf("unexpected content: %s", content)
	}
}
