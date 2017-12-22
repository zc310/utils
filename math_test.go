package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Math(t *testing.T) {
	assert.Equal(t, 0, StrToInt("0"))
	assert.Equal(t, -1, StrToInt("-1"))
	assert.Equal(t, 0, StrToInt(""))
	assert.Equal(t, float32(0), StrToFloat32(""))
	assert.Equal(t, float32(0), Abs32(0))
	assert.Equal(t, float32(1), Abs32(-1))
	assert.Equal(t, float32(1.01), StrToFloat32Def("-", float32(1.01)))
	assert.Equal(t, float32(1.02), StrToFloat32Def("1.02", float32(1.01)))
	assert.Equal(t, 0, StrToIntDef("", 0))
	assert.Equal(t, 123, StrToIntDef("123", 0))
}
