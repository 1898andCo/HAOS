package config_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/config"
)

func TestFuzzyNames(t *testing.T) {
	fn := config.FuzzyNames{}
	fn.AddName("foo", "bar")
	fn.AddName("baz", "qux")
	expected := "qux"
	actual := fn.Names["baz"]
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}

}
