package mocks

import (
	"testing"

	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/spf13/afero"
)

func MockFs(m *testing.M) afero.Fs {
	fs := afero.NewMemMapFs()
	fs.MkdirAll("/etc", 0755)
	fs.MkdirAll("/lib", 0755)
	fs.MkdirAll("/libexec/HAOS", 0755)
	fs.MkdirAll("/sbin", 0755)
	fs.MkdirAll("/share", 0755)
	afero.WriteFile(fs, "/etc/enviromnent", []byte("TEST=1"), 0644)
	fs.MkdirAll("/proc", 0755)
	fs.MkdirAll("/etc/connman/", 0755)
	afero.WriteFile(fs, "/proc/modules", []byte("test"), 0644)
	return fs
}

func NewCloudConfig() *config.CloudConfig {
	var cc = &config.CloudConfig{
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
			Password:  "12345",
			ServerURL: "",
			Token:     "",
			Labels:    map[string]string{},
			K3sArgs:   []string{},
			Environment: map[string]string{
				"TEST1": "1",
				"TEST2": "2",
			},
			Taints: []string{},
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
	return cc

}
