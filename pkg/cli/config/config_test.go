package config

import (
	"testing"

)

func TestConfigCommand(t *testing.T) {
	c := Command()
	expected := "config"
	actual := c.Name
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	expected = "configure HAOS"
	actual = c.Usage
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	expected = "cfg"
	actual = c.ShortName
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
	count := len(c.Flags)
	if count != 5 {
		t.Errorf("Expected 5 flags, got %d", count)
	}
	before := c.Before
	if before == nil {
		t.Errorf("Expected Before to be set")
	}
	action := c.Action
	if action == nil {
		t.Errorf("Expected Action to be set")
	}

}
