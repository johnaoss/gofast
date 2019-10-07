package fs_test

import (
	"log"
	"os"
	"io/ioutil"
	"github.com/johnaoss/speedtest/pkg/fs"
	"testing"
)

func TestCacheWrite(t *testing.T) {
	files, err := fs.New("Test App", "com.johnaoss.speedtest")
	if err != nil {
		t.Fatalf("Failed to initialize app: %v", err)
	}
	cache := files.CacheDir
	name := cache + "filename"

	log.Printf("Creating file: %s\n", name)
	err = ioutil.WriteFile(name, []byte("Hello World!"), os.ModePerm)
	if err != nil {
		t.Errorf("Failed to write cached file: %v", err)
	}

	if err := os.Remove(name); err != nil {
		t.Errorf("Failed to remove file: %v", err)
	}
}
