package cc

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/1898andCo/HAOS/pkg/command"
	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/1898andCo/HAOS/pkg/hostname"
	"github.com/1898andCo/HAOS/pkg/mode"
	"github.com/1898andCo/HAOS/pkg/ssh"
	"github.com/1898andCo/HAOS/pkg/system"
	"github.com/1898andCo/HAOS/pkg/version"
	"github.com/1898andCo/HAOS/pkg/writefile"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"pault.ag/go/modprobe"

	"github.com/spf13/afero"
)

const (
	procModulesFile = "/proc/modules"
	retryMax        = 5
	manifestsDir    = "/var/lib/rancher/k3s/server/manifests"
)

// Syntactic sugar for the bare functions below
func ApplyModules(cfg *config.CloudConfig) error {
	loaded := map[string]bool{}
	f, err := os.Open(procModulesFile)
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		loaded[strings.SplitN(sc.Text(), " ", 2)[0]] = true
	}
	modules := cfg.HAOS.Modules
	for _, m := range modules {
		if loaded[m] {
			continue
		}
		params := strings.SplitN(m, " ", -1)
		logrus.Debugf("module %s with parameters [%s] is loading", m, params)
		if err := modprobe.Load(params[0], strings.Join(params[1:], " ")); err != nil {
			return fmt.Errorf("could not load module %s with parameters [%s], err %v", m, params, err)
		}
		logrus.Debugf("module %s is loaded", m)
	}
	return sc.Err()
}

func ApplySysctls(cfg *config.CloudConfig) error {
	for k, v := range cfg.HAOS.Sysctls {
		elements := []string{"/proc", "sys"}
		elements = append(elements, strings.Split(k, ".")...)
		path := path.Join(elements...)
		if err := afero.WriteFile(system.AppFs, path, []byte(v), 0644); err != nil {
			return err
		}
	}
	return nil
}

func ApplyHostname(cfg *config.CloudConfig) error {
	return hostname.SetHostname(cfg)
}

func ApplyPassword(cfg *config.CloudConfig) error {
	return command.SetPassword(cfg.HAOS.Password)
}

func ApplyRuncmd(cfg *config.CloudConfig) error {
	_, err := command.ExecuteCommand(cfg.Runcmd)
	return err
}

func ApplyBootcmd(cfg *config.CloudConfig) error {
	_, err := command.ExecuteCommand(cfg.Bootcmd)
	return err
}

func ApplyInitcmd(cfg *config.CloudConfig) error {
	_, err := command.ExecuteCommand(cfg.Initcmd)
	return err
}

func ApplyWriteFiles(cfg *config.CloudConfig) error {
	writefile.WriteFiles(cfg)
	return nil
}

// Applys the settings for cloud Config (CC) for manifests
func ApplyBootManifests(cfg *config.CloudConfig) error {
	manifests := cfg.BootManifests
	if len(manifests) == 0 {
		return nil
	}
	filesToWrite := make(map[string][]byte)
	var err error
	for _, m := range manifests {
		var data []byte
		retries := 0
		for retryMax > retries {
			resp, err := http.Get(m.URL)
			if err != nil {
				logrus.Errorf("manifest download failed for %q, retrying [%d/%d]", m.URL, retries, retryMax)
				retries++
				continue
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				logrus.Errorf("manifest download returned non-200 status code for %q, retrying [%d/%d]", m.URL, retries, retryMax)
				retries++
				continue
			}
			data, err = afero.ReadAll(resp.Body)
			if err != nil {
				errors.Wrap(err, "failed reading manifest URL body")
				return err
			}
			if len(data) == 0 {
				return errors.New("empty manifest for %q")
			}
			if m.SHA256 != "" {
				sum := sha256.Sum256(data)
				if fmt.Sprintf("%x", sum) != m.SHA256 {
					return fmt.Errorf("sha256 failed for manifest: %s", m.URL)
				}
			}
			name := m.URL[strings.LastIndex(m.URL, "/")+1:]
			filesToWrite[name] = data
			break
		}
	}
	for file, data := range filesToWrite {
		p := filepath.Join(manifestsDir, file)
		if err := afero.WriteFile(system.AppFs, p, data, 0600); err != nil {
			return err
		}

	}
	return err
}

func ApplySSHKeys(cfg *config.CloudConfig) error {
	return ssh.SetAuthorizedKeys(cfg, false)
}

func ApplySSHKeysWithNet(cfg *config.CloudConfig) error {
	return ssh.SetAuthorizedKeys(cfg, true)
}

func ApplyK3SWithRestart(cfg *config.CloudConfig) error {
	return ApplyK3S(cfg, true, false)
}

func ApplyK3SInstall(cfg *config.CloudConfig) error {
	return ApplyK3S(cfg, true, true)
}

func ApplyK3SNoRestart(cfg *config.CloudConfig) error {
	return ApplyK3S(cfg, false, false)
}

func ApplyK3S(cfg *config.CloudConfig, restart, install bool) error {
	mode, err := mode.Get()
	if err != nil {
		return err
	}
	if mode == "install" {
		return nil
	}

	k3sExists := false
	k3sLocalExists := false
	if _, err := os.Stat("/sbin/k3s"); err == nil {
		k3sExists = true
	}
	if _, err := os.Stat("/usr/local/bin/k3s"); err == nil {
		k3sLocalExists = true
	}

	args := cfg.HAOS.K3sArgs
	vars := []string{
		"INSTALL_K3S_NAME=service",
	}

	if !k3sExists && !restart {
		return nil
	}

	if k3sExists {
		vars = append(vars, "INSTALL_K3S_SKIP_DOWNLOAD=true")
		vars = append(vars, "INSTALL_K3S_BIN_DIR=/sbin")
		vars = append(vars, "INSTALL_K3S_BIN_DIR_READ_ONLY=true")
	} else if k3sLocalExists {
		vars = append(vars, "INSTALL_K3S_SKIP_DOWNLOAD=true")
	} else if !install {
		return nil
	}

	if !restart {
		vars = append(vars, "INSTALL_K3S_SKIP_START=true")
	}

	if cfg.HAOS.ServerURL == "" {
		if len(args) == 0 {
			args = append(args, "server")
		}
	} else {
		vars = append(vars, fmt.Sprintf("K3S_URL=%s", cfg.HAOS.ServerURL))
		if len(args) == 0 {
			args = append(args, "agent")
		}
	}

	if strings.HasPrefix(cfg.HAOS.Token, "K10") {
		vars = append(vars, fmt.Sprintf("K3S_TOKEN=%s", cfg.HAOS.Token))
	} else if cfg.HAOS.Token != "" {
		vars = append(vars, fmt.Sprintf("K3S_CLUSTER_SECRET=%s", cfg.HAOS.Token))
	}

	var labels []string
	for k, v := range cfg.HAOS.Labels {
		labels = append(labels, fmt.Sprintf("%s=%s", k, v))
	}
	if mode != "" {
		labels = append(labels, fmt.Sprintf("haos.io/mode=%s", mode))
	}
	labels = append(labels, fmt.Sprintf("haos.io/version=%s", version.Version))
	sort.Strings(labels)

	for _, l := range labels {
		args = append(args, "--node-label", l)
	}

	for _, taint := range cfg.HAOS.Taints {
		args = append(args, "--kubelet-arg", "register-with-taints="+taint)
	}

	cmd := exec.Command("/usr/libexec/haos/k3s-install.sh", args...)
	cmd.Env = append(os.Environ(), vars...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	logrus.Debugf("Running %s %v %v", cmd.Path, cmd.Args, vars)

	return cmd.Run()
}

func ApplyInstall(cfg *config.CloudConfig) error {
	mode, err := mode.Get()
	if err != nil {
		return err
	}
	if mode != "install" {
		return nil
	}

	cmd := exec.Command("haos", "install")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func ApplyDNS(cfg *config.CloudConfig) error {
	buf := &bytes.Buffer{}
	buf.WriteString("[General]\n")
	buf.WriteString("NetworkInterfaceBlacklist=veth\n")
	buf.WriteString("PreferredTechnologies=ethernet,wifi\n")
	if len(cfg.HAOS.DNSNameservers) > 0 {
		dns := strings.Join(cfg.HAOS.DNSNameservers, ",")
		buf.WriteString("FallbackNameservers=")
		buf.WriteString(dns)
		buf.WriteString("\n")
	} else {
		buf.WriteString("FallbackNameservers=8.8.8.8\n")
	}

	if len(cfg.HAOS.NTPServers) > 0 {
		ntp := strings.Join(cfg.HAOS.NTPServers, ",")
		buf.WriteString("FallbackTimeservers=")
		buf.WriteString(ntp)
		buf.WriteString("\n")
	}

	err := afero.WriteFile(system.AppFs, "/etc/connman/main.conf", buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write /etc/connman/main.conf: %v", err)
	}

	return nil
}

func ApplyWifi(cfg *config.CloudConfig) error {
	if len(cfg.HAOS.Wifi) == 0 {
		return nil
	}

	buf := &bytes.Buffer{}

	buf.WriteString("[WiFi]\n")
	buf.WriteString("Enable=true\n")
	buf.WriteString("Tethering=false\n")

	if buf.Len() > 0 {
		if err := system.AppFs.MkdirAll("/var/lib/connman", 0755); err != nil {
			return fmt.Errorf("failed to mkdir /var/lib/connman: %v", err)
		}
		if err := afero.WriteFile(system.AppFs, "/var/lib/connman/settings", buf.Bytes(), 0644); err != nil {
			return fmt.Errorf("failed to write to /var/lib/connman/settings: %v", err)
		}
	}

	buf = &bytes.Buffer{}

	buf.WriteString("[global]\n")
	buf.WriteString("Name=cloud-config\n")
	buf.WriteString("Description=Services defined in the cloud-config\n")

	for i, w := range cfg.HAOS.Wifi {
		name := fmt.Sprintf("wifi%d", i)
		buf.WriteString("[service_")
		buf.WriteString(name)
		buf.WriteString("]\n")
		buf.WriteString("Type=wifi\n")
		buf.WriteString("Passphrase=")
		buf.WriteString(w.Passphrase)
		buf.WriteString("\n")
		buf.WriteString("Name=")
		buf.WriteString(w.Name)
		buf.WriteString("\n")
	}

	if buf.Len() > 0 {
		return afero.WriteFile(system.AppFs, "/var/lib/connman/cloud-config.config", buf.Bytes(), 0644)
	}

	return nil
}

func ApplyDataSource(cfg *config.CloudConfig) error {
	if len(cfg.HAOS.DataSources) == 0 {
		return nil
	}

	args := strings.Join(cfg.HAOS.DataSources, " ")
	buf := &bytes.Buffer{}

	buf.WriteString("command_args=\"")
	buf.WriteString(args)
	buf.WriteString("\"\n")

	if err := afero.WriteFile(system.AppFs, "/etc/conf.d/cloud-config", buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write to /etc/conf.d/cloud-config: %v", err)
	}

	return nil
}

func ApplyEnvironment(cfg *config.CloudConfig) error {
	if len(cfg.HAOS.Environment) == 0 {
		return nil
	}
	env := make(map[string]string, len(cfg.HAOS.Environment))
	if buf, err := afero.ReadFile(system.AppFs, "/etc/environment"); err == nil {
		scanner := bufio.NewScanner(bytes.NewReader(buf))
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "#") {
				continue
			}
			line = strings.TrimPrefix(line, "export")
			line = strings.TrimSpace(line)
			if len(line) > 1 {
				parts := strings.SplitN(line, "=", 2)
				key := parts[0]
				val := ""
				if len(parts) > 1 {
					if val, err = strconv.Unquote(parts[1]); err != nil {
						val = parts[1]
					}
				}
				env[key] = val
			}
		}
	}
	for key, val := range cfg.HAOS.Environment {
		env[key] = val
	}
	buf := &bytes.Buffer{}
	for key, val := range env {
		buf.WriteString(key)
		buf.WriteString("=")
		buf.WriteString(strconv.Quote(val))
		buf.WriteString("\n")
	}
	if err := afero.WriteFile(system.AppFs, "/etc/environment", buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write to /etc/environment: %v", err)
	}

	return nil
}
