package mode

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/1898andCo/HAOS/pkg/system"

	"github.com/spf13/afero"
)

func Get(cfg *config.CloudConfig, prefix ...string) (string, error) {
	path := filepath.Join(filepath.Join(prefix...), system.StatePath("mode"))
	bytes, err := afero.ReadFile(cfg.Fs, path)
	if os.IsNotExist(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytes)), nil
}
