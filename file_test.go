package utils

import (
	"testing"

	. "gopkg.in/go-playground/assert.v1"
)

func TestExtractFileName(t *testing.T) {
	Equal(t, ExtractFileName("a/b.txt"), "b.txt")
}

func TestChangeFileExt(t *testing.T) {
	Equal(t, ChangeFileExt("a/b.txt", ".ini"), "a/b.ini")
}
