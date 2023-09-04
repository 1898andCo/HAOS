package cc_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/cc"
	"github.com/1898andCo/HAOS/pkg/mocks"
	"github.com/1898andCo/HAOS/pkg/system"
)

func TestMain(m *testing.M) {
	oldFs := system.AppFs
	system.AppFs = mocks.MockFs(m)
	m.Run()
	system.AppFs = oldFs
}

func TestRunApply(t *testing.T) {
	cfg := mocks.NewCloudConfig()
	err := cc.RunApply(cfg, cc.ApplyModules, cc.ApplySysctls, cc.ApplyDNS, cc.ApplyWifi, cc.ApplyEnvironment)
	//cc.ApplyHostname,
	//cc.ApplyPassword,
	//cc.ApplySSHKeysWithNet,
	//cc.ApplyWriteFiles,
	//cc.ApplyBootManifests,
	//cc.ApplyRuncmd,
	//cc.ApplyInstall,
	//cc.ApplyK3SInstall

	if err != nil {
		t.Errorf("RunApply() error = %v", err)
	}
}
