// Package fs provides bundle specific filesystem information.
//
// Eventually should allow me to write and read historic data to and from
// a file.
//
// Reference: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/MacOSXDirectories/MacOSXDirectories.html#//apple_ref/doc/uid/TP40010672-CH10-SW1
package fs

import (
	"os"
)

const (
	support = "~/Library/Application Support/"
	cache   = "~/Library/Caches/"
)

// Eventually should create a file list like the stuff within the CRUX prt-get
// package manager, where it creates a list of files that can be removed on an
// uninstall.

// Would be cool to add an uninstaller for the application on the command line.

// FileSystem is a struct that handles the file system relevant information about
// a macOS specific application.
type FileSystem struct {
	name    string
	id      string
	support string
	cache   string
}

// New returns a FileSystem. `name` should be the human-friendly name of the
// application, and `id` should represent the bundle identifier.
func New(name, id string) FileSystem {
	return FileSystem{name: name, id: id}
}

// CacheDir returns a directory to be used for caching files.
// Will panic if it cannot create the directory that should exist.
// The directory will be in ~/Library/Caches/{ID}
func (f FileSystem) CacheDir() string {
	if f.cache != "" {
		return f.cache
	}
	f.cache = cache + f.id
	if !exists(f.cache) {
		if err := os.Mkdir(f.cache, 755); err != nil {
			panic(err)
		}
	}
	return f.cache
}

// SupportDir returns the directory to be used for long-lasting files.
func (f FileSystem) SupportDir() string {
	if f.support != "" {
		return f.support
	}
	f.support = support + f.name
	if !exists(f.support) {
		if err := os.Mkdir(f.support, 755); err != nil {
			panic(err)
		}
	}
	return f.support
}

// ConfigDir returns the directory to be used for user preferences. This returns
// the MacOS specific verison, and requires me to use the NSUserDefaults class
// which I can'do right now.
func (f FileSystem) ConfigDir() string {
	return ""
}

// exists determines if a directory exists
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
