package writefile_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/1898andCo/HAOS/pkg/mocks"
	"github.com/1898andCo/HAOS/pkg/system"
	"github.com/1898andCo/HAOS/pkg/writefile"
	"github.com/spf13/afero"
)

func TestWriteFiles(t *testing.T) {
	fs := system.AppFs
	defer func() { system.AppFs = fs }()
	cfg := mocks.NewCloudConfig()
	path := "/tmp/test"
	cfg.WriteFiles = []config.File{
		{
			Path:     path,
			Content:  "Zm9vYmFy",
			Encoding: "base64",
		},
	}
	expected := "foobar"
	t.Logf("cfg: %v", cfg.WriteFiles)
	writefile.WriteFiles(cfg)
	// test to see if the files got written to the in-mem filesystem
	content, err := afero.ReadFile(cfg.Fs, path)
	t.Logf("content: %s", content)
	if err != nil {
		t.Errorf("failed to read file: %v", err)
	}
	if string(content) != expected {
		t.Errorf("unexpected content: %s", content)
	}
}
