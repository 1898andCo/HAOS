package util

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/1898andCo/HAOS/pkg/system"
	"github.com/spf13/afero"
)

func WriteFileAtomic(filename string, data []byte, perm os.FileMode) error {
	dir, file := path.Split(filename)
	tempFile, err := afero.TempFile(system.AppFs, dir, fmt.Sprintf(".%s", file))
	if err != nil {
		return err
	}
	defer system.AppFs.Remove(tempFile.Name())
	if _, err := tempFile.Write(data); err != nil {
		return err
	}
	if err := tempFile.Close(); err != nil {
		return err
	}
	if err := system.AppFs.Chmod(tempFile.Name(), perm); err != nil {
		return err
	}
	return system.AppFs.Rename(tempFile.Name(), filename)
}

func HTTPDownloadToFile(url, dest string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := afero.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return WriteFileAtomic(dest, body, 0644)
}

func HTTPLoadBytes(url string) ([]byte, error) {
	var resp *http.Response
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("non-200 http response: %d", resp.StatusCode)
		}

		bytes, err := afero.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return bytes, nil
	}

	return nil, err
}

func ExistsAndExecutable(path string) bool {
	info, err := system.AppFs.Stat(path)
	if err != nil {
		return false
	}

	mode := info.Mode().Perm()
	return mode&os.ModePerm != 0
}

func RunScript(path string, arg ...string) error {
	if !ExistsAndExecutable(path) {
		return nil
	}

	script, err := system.AppFs.Open(path)
	if err != nil {
		return err
	}

	magic := make([]byte, 2)
	if _, err = script.Read(magic); err != nil {
		return err
	}

	cmd := exec.Command("/bin/sh", path)
	if string(magic) == "#!" {
		cmd = exec.Command(path, arg...)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func EnsureDirectoryExists(dir string) error {
	info, err := system.AppFs.Stat(dir)
	if err == nil {
		if !info.IsDir() {
			return fmt.Errorf("%s is not a directory", dir)
		}
	} else {
		err = system.AppFs.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
