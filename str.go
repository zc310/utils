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

// DynIntToCommaText return  CommaText
func DynIntToCommaText(a []int) string {
	var s []string
	for i := 0; i < len(a); i++ {
		s = append(s, strconv.Itoa(a[i]))
	}
	return strings.Join(s, ",")
}
func CommaTextToDynInt(a string) (b []int) {
	var n int
	var err error
	s := strings.Split(a, ",")
	for i := 0; i < len(s); i++ {
		n, err = strconv.Atoi(s[i])
		if err != nil {
			continue
		}
		b = append(b, n)
	}
	return
}



type WideString string

func (p WideString) MarshalJSON() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(string(p))), nil
}
func (p WideString) MarshalText() ([]byte, error) {
	return []byte(strconv.QuoteToASCII(string(p))), nil
}
