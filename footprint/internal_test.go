package footprint

import (
	"testing"
	"bytes"
)

// TestWriteLine is a sanity check to see if the writes are done properly.
func TestWriteLine(t *testing.T) {
	const expected = "-rw-------\troot/root\t/tmp/gofast/test\n"
	b := new(bytes.Buffer)
	writeLine(b, "-rw-------", "root/root", "/tmp/gofast/test")
	if b.String() != expected {
		t.Error("Given:", b.String(), "Expected:", expected)
	}	
}
