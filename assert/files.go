package assert

import (
	"os"
	"path/filepath"
	"testing"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

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

func Open(t *testing.T, filename string) *os.File {
	f, err := os.Open(filename)
	IsNil(t, err)
	return f
}
