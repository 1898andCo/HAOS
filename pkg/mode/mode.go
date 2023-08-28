package mode

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/1898andCo/HAOS/pkg/system"
	
	"github.com/spf13/afero"
)

func Get(prefix ...string) (string, error) {
	bytes, err := afero.ReadFile(system.AppFs, filepath.Join(filepath.Join(prefix...), system.StatePath("mode")))
	if os.IsNotExist(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytes)), nil
}
