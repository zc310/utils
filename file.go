package utils

import (
	"fmt"
	"path/filepath"
	"os"
)

// ExtractFileName  extracts the file name substring.
func ExtractFileName(s string) string {
	_, name := filepath.Split(s)
	return name
}

// ChangeFileExt  changes the Extension value of a file FileName, returning the new value as a string.
func ChangeFileExt(s, ext string) string {
	return fmt.Sprintf("%s%s", s[:len(s)-len(filepath.Ext(s))], ext)

}

// IsDirectory check file or directory
func IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	return fileInfo.IsDir(), err
}