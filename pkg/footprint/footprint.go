// Package footprint maintains a footprint of all the files used in an
// application. 
//
// This allows centralized management of files, and eventually will allow a method
// to remove all files associated with this program from a user's disk. Essentially
// acting as a complete uninstall, instead of leaving leftover files.
package footprint

import (
	"io"
	"bytes"
	"path/filepath"
	"sync"
	"fmt"
	"os"
)	

const (
	// default perms
	perms = 0744
)

// Footprint represents a list of files that a program may use.
// This allows the central handling of all temporary files a program may create.
type Footprint struct {
	// filename is the name of the file this is backed by.
	// this file should not be read by others.
	filename string

	// file is the underlying file, may be nil depending on if this writes to
	// disk or not.
	// todo: determine how exactly to effectively sync the contents of this struct
	// to the disk. this is a hard problem.
	file *os.File

	// mutex to allow concurrent access
	mu sync.RWMutex

	// files is a list of all the files currently being watched in the footprint
	files []File
}

// New allocates a new footprint file by creating the file with the given
// filename and directory.
func New(dir, filename string) (*Footprint, error) {
	file, err := os.Create(filepath.Join(dir, filename))
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	return &Footprint{
		filename: file.Name(),
		file: file,
		files: make([]File, 0),
	}, nil
}

// Parse reads from a file and allocates a Footprint representing the info
// in the read file.
// TODO: Do this.
func Parse(file *os.File) (*Footprint, error) {
	return nil, nil
}

// String writes out the expected contents of the file to a string.
func (f *Footprint) String() string {
	buf := new(bytes.Buffer)
	if err := f.write(buf); err != nil {
		return err.Error()
	}
	return buf.String()
}

// Name gets the name of the underlying file.
// Will return the empty string if the underlying file doesn't exist.
func (f *Footprint) Name() string {
	if f == nil || f.file == nil {
		return ""
	}
	return f.file.Name()
}

// Len returns the current number of files stored within the footprint.
func (f *Footprint) Len() int {
	f.mu.RLock()
	num := len(f.files)
	f.mu.RUnlock()
	return num
}

// Add registers a file with the footprint.
func (f *Footprint) Add(file *os.File) error {
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("could not stat file: %w", err)
	}

	// Most likely can't sync the file here.
	f.mu.Lock()
	f.files = append(f.files, NewFile(file.Name(), info))
	f.mu.Unlock()

	return nil
}

// Close closes the footprint file.
func (f *Footprint) Close() error {
	file, err := f.getfile()
	if err != nil {
		return err
	}

	return f.write(file)
}

// getfile gets the underlying footprint file, or creates one if none exists.
func (f *Footprint) getfile() (*os.File, error) {
	if f.file != nil {
		return f.file, nil
	}

	var err error
	f.file, err = os.OpenFile(f.filename, os.O_RDWR, perms)
	if err != nil {
		return nil, fmt.Errorf("failed to open footprint file: %w", err)
	}

	return f.file, nil
}

// write outputs the representation of the file to an io.Writer.
func (f *Footprint) write(w io.Writer) (error) {
	b := new(bytes.Buffer)
	for _, elem := range f.files {
		writeLine(b, elem.Perms().String(), "root/root", elem.Path())
		if _, err := w.Write(b.Bytes()); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
		b.Reset()
	}

	return nil
}

// writeline writes a tab-spaced line of text given a varadic number of inputs to
// a given buffer. after all inputs are written, a newline will be appended.
func writeLine(b *bytes.Buffer, inputs ...string) {
	var size int 
	for i := range inputs {
		size += len(inputs[i])
	}

	if size > b.Cap() - b.Len() {
		b.Grow(size - b.Cap() + b.Len())
	}

	for i := range inputs {
		b.WriteString(inputs[i])
		if i != len(inputs)-1 {
			b.WriteByte('\t')
		}
	}
	b.WriteByte('\n')
}