// Package fs provides bundle specific filesystem information.
//
// Eventually should allow me to write and read historic data to and from
// a file.
//
// Reference: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/MacOSXDirectories/MacOSXDirectories.html#//apple_ref/doc/uid/TP40010672-CH10-SW1
package fs

import (
	"fmt"
	"os"
)

const (
	support = "/Library/Application Support/"
	cache   = "/Library/Caches/"
)

// Eventually should create a file list like the stuff within the CRUX prt-get
// package manager, where it creates a list of files that can be removed on an
// uninstall.

// Would be cool to add an uninstaller for the application on the command line.

// FileSystem is a struct that handles the file system relevant information about
// a macOS specific application.
type FileSystem struct {
	// name the human-friendly name of the application
	name    string
	// id is the bundle identifier of the application. 
	id      string
	// SupportDirDir is the filepath of the Application SupportDir directory, 
	// used for long-lasting files.
	SupportDir string
	// CacheDir is the filepath of the directory to be used for caching files.
	CacheDir   string
	// HomeDir is the home directory of the user
	HomeDir string
}

// New returns a FileSystem. `name` should be the human-friendly name of the
// application, and `id` should represent the bundle identifier.
func New(name, id string) (*FileSystem, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to obtain user home dir: %w", err)
	}

	f := FileSystem{
		name: name,
		id: id,
		SupportDir: home + support + name + "/",
		CacheDir: home + cache + id + "/",
		HomeDir: home,
	}

	if !exists(f.CacheDir) {
		if err := os.MkdirAll(f.CacheDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to make cache dir: %w", err)
		}
	}

	if !exists(f.SupportDir) {
		if err := os.MkdirAll(f.SupportDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to make support dir: %w", err)
		}
	}

	return &f, nil
}

func (f FileSystem) TempFile(name string) (*os.File, error) {
	filename := f.CacheDir + name
	// TODO: Add to main list
	return os.Create(filename)
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