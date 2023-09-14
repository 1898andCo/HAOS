package cli_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/cli"
)

func TestRCCommand(t *testing.T) {
	commandResult := cli.RCCommand()
	expected := "rc"
	actual := commandResult.Name
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	expected =  "early phase \"run commands\" / \"run control\""
	actual = commandResult.Usage
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	expectedCount := 0
	actualCount := len(commandResult.Flags)
	if expectedCount != actualCount {
		t.Errorf("Expected %d flags, got %d", expectedCount, actualCount)
	}
	if commandResult.Before == nil {
		t.Errorf("Expected Before not to be nil")
	}
	if commandResult.Action == nil {
		t.Errorf("Expected Action not to be nil")
	}

}
