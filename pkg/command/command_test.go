package command_test

import (
	"context"
)

type MockShell struct {
	// an output and error to be returned when command is executed
	Output []byte
	Err    error
	// store the last executed command
	// in case we need to test what command string did the code produce
	LastCommand string
}

func (t *MockShell) Execute(ctx context.Context, cmd string) ([]byte, error) {
	t.LastCommand = cmd

	return t.Output, t.Err
}
