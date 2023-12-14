package cliinstall

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/1898andCo/HAOS/pkg/config"
	"github.com/1898andCo/HAOS/pkg/questions"
	"github.com/ghodss/yaml"
)

func Run() error {
	fmt.Println("\nRunning HAOS configuration")

	cfg, err := config.ReadConfig()
	if err != nil {
		return err
	}

	isInstall, err := Ask(&cfg)
	if err != nil {
		return err
	}

	if isInstall {
		return runInstall(cfg)
	}

	bytes, err := config.ToBytes(cfg)
	if err != nil {
		return err
	}

	f, err := os.Create(config.SystemConfig)
	if err != nil {
		f, err = os.Create(config.LocalConfig)
		if err != nil {
			return err
		}
	}
	defer f.Close()

	if _, err := f.Write(bytes); err != nil {
		return err
	}

	f.Close()
	return runCCApply()
}

func runCCApply() error {
	cmd := exec.Command(os.Args[0], "config", "--install")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func runInstall(cfg config.CloudConfig) error {
	var (
		err      error
		tempFile *os.File
	)

	installBytes, err := config.PrintInstall(cfg)
	if err != nil {
		return err
	}

	if !cfg.HAOS.Install.Silent {
		val, err := questions.PromptBool("\nConfiguration\n"+"-------------\n\n"+
			string(installBytes)+
			"\nYour disk will be formatted and HAOS will be installed with the above configuration.\nContinue?", false)
		if err != nil || !val {
			return err
		}
	}

	if cfg.HAOS.Install.ConfigURL == "" {
		tempFile, err = ioutil.TempFile("/tmp", "haos.XXXXXXXX")
		if err != nil {
			return err
		}
		defer tempFile.Close()

		cfg.HAOS.Install.ConfigURL = tempFile.Name()
	}

	ev, err := config.ToEnv(cfg)
	if err != nil {
		return err
	}

	if tempFile != nil {
		cfg.HAOS.Install = nil
		bytes, err := yaml.Marshal(&cfg)
		if err != nil {
			return err
		}
		if _, err := tempFile.Write(bytes); err != nil {
			return err
		}
		if err := tempFile.Close(); err != nil {
			return err
		}
		defer os.Remove(tempFile.Name())
	}

	cmd := exec.Command("/usr/libexec/haos/install")
	cmd.Env = append(os.Environ(), ev...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
