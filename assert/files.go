package assert

import (
	"os"
	"path/filepath"
	"testing"
)

// FileExists checks that a file exists with the given filename.
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Exists checks that each of the files exist in the root directory.
func Exists(t *testing.T, root string, files ...string) {
	for _, file := range files {
		dir := filepath.Join(root, file)
		if FileExists(dir) {
			continue
		}
		t.Logf("File '%s' does not exist.", dir)
		t.Fail()
	}
}

// RemoveAll removes all the files listed in the given root directory.
func RemoveAll(root string, files ...string) {
	for _, file := range files {
		dir := filepath.Join(root, file)
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			continue
		}
		os.RemoveAll(dir)
	}
}

// Open opens the given file and fails if the Open call produces an error.
func Open(t *testing.T, filename string) *os.File {
	f, err := os.Open(filename)
	IsNil(t, err)
	return f
}
