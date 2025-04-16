package utils

import (
	"io"
	"os"
)

func IsDir(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		return false
	}
	if info.IsDir() {
		return true
	}
	return false
}

func CopyFile(source string, dest string, sync bool) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}
	if sync {
		err = destFile.Sync()
		if err != nil {
			return err
		}
	}
	return nil
}
