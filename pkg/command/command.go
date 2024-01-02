package command

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

type abstract interface {
	Command(name string, arg ...string) *exec.Cmd
}

type concrete struct{}

func (concrete) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command(name, arg...)
}

var impl abstract = concrete{}

func ExecuteCommand(commands []string) error {
	for _, cmd := range commands {
		logrus.Debugf("running cmd `%s`", cmd)
		ce := impl.Command("sh", "-c", cmd)
		ce.Stdout = os.Stdout
		ce.Stderr = os.Stderr
		if err := ce.Run(); err != nil {
			return fmt.Errorf("failed to run %s: %v", cmd, err)
		}
	}
	return nil
}

func SetPassword(password string) error {
	if password == "" {
		return nil
	}
	cmd := impl.Command("chpasswd")
	if strings.HasPrefix(password, "$") {
		cmd.Args = append(cmd.Args, "-e")
	}
	// TODO(username): username should be hardcoded to 1898andco
	cmd.Stdin = strings.NewReader(fmt.Sprint("rancher:", password))
	cmd.Stdout = os.Stdout
	errBuffer := &bytes.Buffer{}
	cmd.Stderr = errBuffer
	err := cmd.Run()
	if err != nil {
		os.Stderr.Write(errBuffer.Bytes())
	}
	return err
}
