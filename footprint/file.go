package footprint

import (
	"syscall"
	"os"
)

// File represents our needed info for an entry in our footprint file.
type File interface {
	Path() string
	Perms() os.FileMode
}

// file is our implementation of the File interface.
type file struct {
	info os.FileInfo
	path string 
	uid uint32
	gid uint32
}

// NewFile returns a new file that satisifes the File interface.
func NewFile(path string, info os.FileInfo) File {
	return &file{
		path: path,
		info: info,
		uid: info.Sys().(*syscall.Stat_t).Uid,
		gid: info.Sys().(*syscall.Stat_t).Gid,
	}
}

// Path returns the absolute path of the file.
func (f *file) Path() string {
	return f.path
}

// Perms returns the permissions of the file.
func (f *file) Perms() os.FileMode {
	return f.info.Mode()
}


