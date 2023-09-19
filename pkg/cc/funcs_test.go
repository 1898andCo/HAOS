package cc_test

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/cc"
	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/1898andCo/HAOS/pkg/mocks"
	"github.com/spf13/afero"
)

var cconfig = mocks.NewCloudConfig()

func testFunc(t *testing.T, f func(cfg *config.CloudConfig) error, label string) {
	err := f(cconfig)
	if err != nil {
		t.Errorf("%s error = %v, wantErr %v", label, err, false)
	}
}

func TestApplyModules(t *testing.T) {

	testFunc(t, cc.ApplyModules, "ApplyModules()")
}

// TODO: this is not very good, it skips the actual code as it
// doesn't iterate over anything due to missing values in config
func TestApplySysctls(t *testing.T) {
	testFunc(t, cc.ApplySysctls, "ApplySysctls()")
}

func TestApplyHostname(t *testing.T) {
	err := cc.ApplyHostname(mocks.NewCloudConfig())
	if err == nil {
		t.Error("Expected operation not permitted error, got nil")
	}
}

// bad test, only tests a crappy fail case
func TestApplyPassword(t *testing.T) {
	cfg := mocks.NewCloudConfig()
	err := cc.ApplyPassword(cfg)
	afero.WriteFile(cfg.Fs, "/etc/passwd", []byte("root:x:0:0:root:/root:/bin/bash"), 0644)
	if err == nil {
		t.Error("Expected Authentication token manipulation error, got nil")
	}
}

func TestApplyDNS(t *testing.T) {

	testFunc(t, cc.ApplyDNS, "ApplyDNS()")
}

func TestApplyWifi(t *testing.T) {
	testFunc(t, cc.ApplyWifi, "ApplyWifi()")
}

func TestApplyRuncmd(t *testing.T) {
	testFunc(t, cc.ApplyRuncmd, "ApplyRuncmd()")
}

func TestApplyBootcmd(t *testing.T) {
	testFunc(t, cc.ApplyBootcmd, "ApplyBootcmd()")
}

func TestApplyInitcmd(t *testing.T) {
	testFunc(t, cc.ApplyInitcmd, "ApplyInitcmd()")
}

func TestApplyWriteFiles(t *testing.T) {
	testFunc(t, cc.ApplyWriteFiles, "ApplyWriteFiles()")
}

func TestApplyBootManifests(t *testing.T) {
	testFunc(t, cc.ApplyBootManifests, "ApplyBootManifests()")
}

func TestApplyEnviromnent(t *testing.T) {
	testFunc(t, cc.ApplyEnvironment, "ApplyEnvironment()")
}

func TestApplyInstall(t *testing.T) {
	testFunc(t, cc.ApplyInstall, "ApplyInstall()")
}