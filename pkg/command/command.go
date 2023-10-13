package command

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"context"

	"github.com/sirupsen/logrus"
)

// Wrap the exec call in a struct so that we can mock it in our tests.
type Shell interface {
	// Execute runs the given command and returns the output and error.
	Execute(ctx context.Context, cmd string) (output []byte, err error)
}

// LocalShell is the default shell that executes commands locally, basically
// calling `sh -c <cmd>` at the end of the day.
type LocalShell struct{}

func (LocalShell) Execute(ctx context.Context, cmd string) ([]byte, error) {
	wrapperCmd := exec.CommandContext(ctx, "sh", "-c", cmd)
	return wrapperCmd.CombinedOutput()
}

// DefaultShell will be overridden in tests
// Set the default shell to be the local shell.
// TODO: still not happy with this pattern
var DefaultShell Shell = LocalShell{}

func ExecuteCommand(commands []string) ([]string, error) {
	var results []string
	for _, cmd := range commands {
		logrus.Debugf("running cmd `%s`", cmd)
		out, err := DefaultShell.Execute(context.Background(), cmd)
		if err != nil {
			return results, fmt.Errorf("failed to run %s: %v", cmd, err)
		}
		results = append(results, string(out))
	}
	return results, nil
}

func SetPassword(password string) error {
	if password == "" {
		return nil
	}
	cmd := exec.Command("chpasswd")
	if strings.HasPrefix(password, "$") {
		cmd.Args = append(cmd.Args, "-e")
	}
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
