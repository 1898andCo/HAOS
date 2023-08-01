// Package system abstracts the filesystem layout of the system and exposes functions to copy and move files
package system

import "path/filepath"

const (
	// DefaultRootDir represents where persistent installations are located
	DefaultRootDir = "/HAOS/system"
	// DefaultDataDir represents where persistent state is located
	DefaultDataDir = "/HAOS/data"
	// DefaultLocalDir represents where local, persistent configuration is located
	DefaultLocalDir = "/var/lib/1898andCo/k3os"
	// DefaultStateDir represents where ephemeral state is located
	DefaultStateDir = "/run/HAOS"
)

var (
	rootDirectory  = DefaultRootDir
	dataDirectory  = DefaultDataDir
	localDirectory = DefaultLocalDir
	stateDirectory = DefaultStateDir
)

// RootPath joins any number of elements into a single path underneath the persistent installation root, by default `DefaultRootDir`
func RootPath(elem ...string) string {
	return filepath.Join(rootDirectory, filepath.Join(elem...))
}

// DataPath joins any number of elements into a single path underneath the persistent state root, by default `DefaultDataDir`
func DataPath(elem ...string) string {
	return filepath.Join(dataDirectory, filepath.Join(elem...))
}

// LocalPath joins any number of elements into a single path underneath the persistent configuration root, by default `DefaultLocalDir`
func LocalPath(elem ...string) string {
	return filepath.Join(localDirectory, filepath.Join(elem...))
}

// StatePath joins any number of elements into a single path underneath the ephemeral state root, by default `DefaultStateDir`
func StatePath(elem ...string) string {
	return filepath.Join(stateDirectory, filepath.Join(elem...))
}
