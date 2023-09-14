package config

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestDataSource(t *testing.T) {
	cc, err := readersToObject(func() (map[string]interface{}, error) {
		return map[string]interface{}{
			"HAOS": map[string]interface{}{
				"Datasources": []string{"foo"},
			},
		}, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	logrus.Info(cc.HAOS)
	if len(cc.HAOS.DataSources) != 1 {
		t.Fatal("no datasources")
	}
	if cc.HAOS.DataSources[0] != "foo" {
		t.Fatalf("%s != foo", cc.HAOS.DataSources[0])
	}
}

func TestAuthorizedKeys(t *testing.T) {
	c1 := map[string]interface{}{
		"SSHAuthorizedKeys": []string{
			"one...",
		},
	}
	c2 := map[string]interface{}{
		"SSHAuthorizedKeys": []string{
			"two...",
		},
	}
	cc, err := readersToObject(
		func() (map[string]interface{}, error) {
			return c1, nil
		},
		func() (map[string]interface{}, error) {
			return c2, nil
		},
	)
	logrus.Infof("%+v", cc)
	if len(cc.SSHAuthorizedKeys) != 1 {
		t.Fatal(err, fmt.Sprintf("got %d keys, expected 2", len(cc.SSHAuthorizedKeys)))
	}
	if err != nil {
		t.Fatal(err)
	}
}
