package main

// Copyright 2023 1898andCo, Inc
// SPDX-License-Identifier: Apache-2.0

import (
	"os"
	"path/filepath"

	"github.com/1898andCo/HAOS/pkg/cli/app"
	"github.com/1898andCo/HAOS/pkg/enterchroot"
	"github.com/1898andCo/HAOS/pkg/transferroot"
	"github.com/docker/docker/pkg/mount"
	"github.com/docker/docker/pkg/reexec"
	"github.com/sirupsen/logrus"
)

func main() {
	reexec.Register("/init", initrd)      // mode=live
	reexec.Register("/sbin/init", initrd) // mode=local
	reexec.Register("enter-root", enterchroot.Enter)

	if !reexec.Init() {
		app := app.New()
		args := []string{app.Name}
		path := filepath.Base(os.Args[0])
		if path != app.Name && app.Command(path) != nil {
			args = append(args, path)
		}
		args = append(args, os.Args[1:]...)
		// this will bomb if the app has any non-defaulted, required flags
		err := app.Run(args)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}

func initrd() {
	enterchroot.DebugCmdline = "haos.debug"
	transferroot.Relocate()
	if err := mount.Mount("", "/", "none", "rw,remount"); err != nil {
		logrus.Errorf("failed to remount root as rw: %v", err)
	}
	if err := enterchroot.Mount("./haos/data", os.Args, os.Stdout, os.Stderr); err != nil {
		logrus.Fatalf("failed to enter root: %v", err)
	}
}
