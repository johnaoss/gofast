package footprint_test

import (
	"io/ioutil"
	"testing"
	"os"

	"github.com/johnaoss/speedtest/pkg/footprint"
)

var (
	dir = os.TempDir()
)

// TestNew tests if we can properly init and then add a file to the footprint
func TestNew(t *testing.T) {
	// Initialize new footprint file
	fp, err := footprint.New(dir, "testnew.footprint")
	if err != nil {
		t.Errorf("failed to init footprint: %v", err)
	}
	t.Log("New Footprint:", fp.Name())
	defer os.Remove(fp.Name())

	// Detect if after adding a new file, we can see that change reflected in
	// the footprint.
	exFile := newfile(t)
	defer os.Remove(exFile.Name())
	if err := fp.Add(exFile); err != nil {
		t.Errorf("failed to add file: %v", err)
	}
	if fp.Len() != 1 {
		t.Errorf("length should be 1, instead was: %d", fp.Len())
	}

	// TODO: Sync to disk.
}

// newfile is a helper that creates a temporary file.
func newfile(t *testing.T) *os.File {
	t.Helper()
	f, err := ioutil.TempFile(dir, "file")
	if err != nil {
		t.Errorf("Failed to create temp file: %v", err)
	}

	t.Log("Created file:", f.Name())
	return f
}