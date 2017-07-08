package utils

import (
	. "gopkg.in/go-playground/assert.v1"
	"testing"
)

func Test_Math(t *testing.T) {
	Equal(t, 0, StrToInt("0"))
	Equal(t, -1, StrToInt("-1"))
	Equal(t, 0, StrToInt(""))
	Equal(t, float32(0), StrToFloat32(""))
	Equal(t, float32(0), Abs32(0))
	Equal(t, float32(1), Abs32(-1))
	Equal(t, float32(1.01), StrToFloat32Def("-", float32(1.01)))
	Equal(t, float32(1.02), StrToFloat32Def("1.02", float32(1.01)))
	Equal(t, 0, StrToIntDef("", 0))
	Equal(t, 123, StrToIntDef("123", 0))
}
