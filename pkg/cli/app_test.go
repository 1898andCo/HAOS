package cli_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/cli"
	"github.com/1898andCo/HAOS/pkg/version"
)

func TestNewApp(t *testing.T) {
	app := cli.NewApp()
	expected := "haos"
	actual := app.Name
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	expected = "Booting to k3s so you don't have to"
	actual = app.Usage
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	expected = version.Version
	actual = app.Version
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	count := len(app.Commands)
	if count != 4 {
		t.Errorf("Expected 4 commands, got %d", count)
	}

}
