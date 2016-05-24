package assert

import (
	"os"
	"path/filepath"
	"testing"
)

func Exists(t *testing.T, root string, files ...string) {
	for _, file := range files {
		dir := filepath.Join(root, file)
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			t.Fail()
		}
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
