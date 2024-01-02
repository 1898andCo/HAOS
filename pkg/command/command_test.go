package command

import (
	"os/exec"
	"testing"
)

type Mock struct{}

func (Mock) Command(name string, arg ...string) *exec.Cmd {
	return exec.Command("echo", "mock")
}

func TestExecuteCommand(t *testing.T) {
	impl = Mock{}
	err := ExecuteCommand([]string{"echo"})
	if err != nil {
		t.Errorf("failed to execute command: %v", err)
	}
}

func TestSetPassword(t *testing.T) {
	impl = Mock{}
	err := SetPassword("mock")
	if err != nil {
		t.Errorf("failed to set password: %v", err)
	}
}
