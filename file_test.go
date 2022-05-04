package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractFileName(t *testing.T) {
	assert.Equal(t, ExtractFileName("a/b.txt"), "b.txt")
}

func TestChangeFileExt(t *testing.T) {
	assert.Equal(t, ChangeFileExt("a/b.txt", ".ini"), "a/b.ini")
}
func TestIsDirectory(t *testing.T) {
	ok, err := IsDirectory("/tmp")
	assert.Equal(t, err, nil)
	assert.Equal(t, ok, true)
}
