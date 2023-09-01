package cc_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/cc"
	"github.com/1898andCo/HAOS/pkg/mocks"
	"github.com/1898andCo/HAOS/pkg/system"
	"github.com/spf13/afero"
)

func TestRunApply(t *testing.T) {
	system.AppFs = afero.NewMemMapFs()
	system.AppFs.MkdirAll("/etc", 0755)
	afero.WriteFile(system.AppFs,"/etc/enviromnent", []byte("TEST=1"), 0644)
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
