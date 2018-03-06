package utils

import (
	"strconv"
	"strings"
)

// IfEmpty return first non-empty string
func IfEmpty(a, b string) string {
	if a == "" {
		return b
	}
	return a
}

// DynArrayIntToCommaText return  CommaText
func DynArrayIntToCommaText(a []int) string {
	var s []string
	for i := 0; i < len(a); i++ {
		s = append(s, strconv.Itoa(a[i]))
	}
	return strings.Join(s, ",")
}

type WideString string

func (p WideString) MarshalJSON() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(string(p))), nil
}
func (p WideString) MarshalText() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(string(p))), nil
}
