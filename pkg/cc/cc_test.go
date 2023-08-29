package cc_test

import (
	"fmt"
	"testing"

	"github.com/1898andCo/HAOS/pkg/cc"
	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/1898andCo/HAOS/pkg/system"
	"github.com/spf13/afero"
)

/*
	ApplyModules,
	ApplySysctls,
	ApplyHostname,
	ApplyDNS,
	ApplyWifi,
	ApplyPassword,
	ApplySSHKeysWithNet,
	ApplyWriteFiles,
	ApplyBootManifests,
	ApplyEnvironment,
	ApplyRuncmd,
	ApplyInstall,
	ApplyK3SInstall,
*/

var Cconfig config.CloudConfig = config.CloudConfig{
	Runcmd:            []string{"echo 'runcmd test'"},
	Bootcmd:           []string{"echo 'bootcmd test'"},
	Initcmd:           []string{"echo 'initcmd test'"},
	SSHAuthorizedKeys: []string{},
	BootManifests: []config.Manifest{
		{
			SHA256: "DEADBEEF",
			URL:    "",
		},
	},
	WriteFiles: []config.File{
		{
			Content:            "testcontent",
			Encoding:           "testencoding",
			Owner:              "testowner",
			Path:               "testpath",
			RawFilePermissions: "0644",
		},
	},
	Hostname: "testhostname",
	HAOS: config.HAOS{
		DataSources: []string{},
		Modules:     []string{},
		Sysctls:     map[string]string{},
		NTPServers: []string{
			"0.pool.ntp.org",
		},
		DNSNameservers: []string{
			"8.8.8.8",
		},
		Wifi: []config.Wifi{
			{
				Name:       "testwifi",
				Passphrase: "testpassphrase",
			},
		},
		Password:    "12345",
		ServerURL:   "",
		Token:       "",
		Labels:      map[string]string{},
		K3sArgs:     []string{},
		Environment: map[string]string{},
		Taints:      []string{},
		Install: &config.Install{
			ForceEFI:  false,
			Device:    "",
			ConfigURL: "",
			Silent:    false,
			ISOURL:    "",
			PowerOff:  false,
			NoFormat:  false,
			Debug:     false,
			TTY:       "",
		},
	},
}

func testFunc(t *testing.T, f func(cfg *config.CloudConfig) error, label string) {
	err := f(&Cconfig)
	if err != nil {
		if fmt.Sprintf("%s", err) == "operation not permitted" {
			t.Logf("%s error = %v, wantErr %v", label, err, false)
		} else {
			t.Errorf("%s error = %v, wantErr %v", label, err, false)
		}
	}
}

func setupFS() {
	system.AppFs = afero.NewMemMapFs()
}

func TestApplyModules(t *testing.T) {
	setupFS()
	system.AppFs.MkdirAll("/proc", 0755)
	afero.WriteFile(system.AppFs, "/proc/modules", []byte("test"), 0644)
	testFunc(t, cc.ApplyModules, "ApplyModules()")
}

// TODO: this is not very good, it skips the actual code as it
// doesn't iterate over anything due to missing values in config
func TestApplySysctls(t *testing.T) {
	testFunc(t, cc.ApplySysctls, "ApplySysctls()")
}

func TestApplyHostname(t *testing.T) {
	testFunc(t, cc.ApplyHostname, "ApplyHostname()")
}

func TestApplyPassword(t *testing.T) {
	setupFS()
	system.AppFs.MkdirAll("/etc", 0755)
	afero.WriteFile(system.AppFs, "/etc/passwd", []byte("root:x:0:0:root:/root:/bin/bash"), 0644)
	testFunc(t, cc.ApplyPassword, "ApplyPassword()")
}

func TestApplyDNS(t *testing.T) {
	setupFS()
	system.AppFs.MkdirAll("/etc/connman/", 0755)
	testFunc(t, cc.ApplyDNS, "ApplyDNS()")
}

func TestApplyWifi(t *testing.T) {
	setupFS()
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
