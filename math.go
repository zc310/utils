package utils

import "strconv"

//Abs32 Gives the absolute value of a float32
func Abs32(x float32) float32 {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

//StrToFloat32 Convert a float32 string into a float32 value
func StrToFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0
	}
	return float32(f)
}

//StrToFloat32Def  Convert a string into an float32 value with default
func StrToFloat32Def(s string, def float32) float32 {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return def
	}
	return float32(f)
}

//StrToInt Convert an int string into an int value
func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

//StrToIntDef Convert a string into an int value with default
func StrToIntDef(s string, def int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return i
}
