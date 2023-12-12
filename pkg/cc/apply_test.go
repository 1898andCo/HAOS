package cc_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/cc"
	"github.com/1898andCo/HAOS/pkg/mocks"
)

func TestRunApply(t *testing.T) {
	cfg := mocks.NewCloudConfig()
	err := cc.RunApply(cfg, cc.ApplyModules, cc.ApplySysctls, cc.ApplyDNS, cc.ApplyWifi, cc.ApplyEnvironment)

	if err != nil {
		t.Errorf("RunApply() error = %v", err)
	}
}
