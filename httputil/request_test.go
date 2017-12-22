package httputil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetString(t *testing.T) {
	_, err := GetString("http://httpbin.org/gzip")
	assert.Equal(t, err, nil)
}
