package command_test

import (
	"context"
	"testing"

	"github.com/1898andCo/HAOS/pkg/command"
)

type MockShell struct {
	// an output and error to be returned when command is executed
	Output []byte
	Err    error
}

func (t *MockShell) Execute(ctx context.Context, cmd string) ([]byte, error) {
	return t.Output, t.Err
}

func TestExecuteCommand(t *testing.T) {
	command.DefaultShell = &MockShell{
		Output: []byte("hello"),
		Err:    nil,
	}

	results, err := command.ExecuteCommand([]string{"echo hello"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(results) != 1 {
		t.Errorf("unexpected result: %v", results)
	}
	if results[0] != "hello" {
		t.Errorf("unexpected result: %v", results)
	}

}
