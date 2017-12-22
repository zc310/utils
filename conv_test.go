package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	assert.Equal(t, 0, GetInt(0))
	assert.Equal(t, 1, GetInt("1"))
	assert.Equal(t, 1, GetInt(" 1 "))

	assert.Equal(t, "1", GetString(1))
	assert.Equal(t, "1", GetString("1"))

	assert.Equal(t, true, GetBool(true))
	assert.Equal(t, true, GetBool("true"))
	assert.Equal(t, true, GetBool("1"))
	assert.Equal(t, false, GetBool("0"))
	assert.Equal(t, false, GetBool(0))
}
