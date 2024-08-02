package fileutil

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

func Ext(file string) string {
	return path.Ext(file)
}

func NameWithoutExt(name string) string {
	name = filepath.Base(name)
	return strings.TrimSuffix(name, filepath.Ext(name))
}

func Rename(file string, to string) error {
	return os.Rename(file, to)
}

func IsFile(filePath string) bool {
	f, e := os.Stat(filePath)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// CreateFileIfNotExists creates file specified if not exists
func CreateFileIfNotExists(name string) error {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		// create
		return CreateRecursively(name)
	} else if os.IsExist(err) {
		// already exists
		return nil
	}
	return err
}

// CreateRecursively creates file recursively.
func CreateRecursively(name string) error {
	if !strings.Contains(name, "/") {
		// just a single filename
		_, err := os.Create(name)
		return err
	}

	i := strings.LastIndex(name, "/")
	path := name[:i]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}

	_, err := os.Create(name)
	return err
}

func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
