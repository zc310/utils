package floatf

import (
	"strconv"
	"strings"
)

func Join(a []float64, sep string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return strconv.FormatFloat(a[0], 'f', -1, 64)
	}

	var b strings.Builder
	b.WriteString(strconv.FormatFloat(a[0], 'f', -1, 64))
	for j := 1; j < len(a); j++ {
		b.WriteString(sep)
		b.WriteString(strconv.FormatFloat(a[j], 'f', -1, 64))
	}

	return b.String()
}
func Parse(a string) ([]float64, error) {
	var (
		err error
		ta  []float64
		t   float64
		b   strings.Builder
	)

	for _, r := range a {
		// . 0~9
		if (r == 46) || ('0' <= r && r <= '9') || r == '-' {
			b.WriteRune(r)
		} else {
			if b.Len() > 0 {
				t, err = strconv.ParseFloat(b.String(), 64)
				if err != nil {
					return nil, err
				}
				ta = append(ta, t)
				b.Reset()
			}
		}
	}
	if b.Len() > 0 {
		t, err = strconv.ParseFloat(b.String(), 64)
		if err != nil {
			return nil, err
		}
		ta = append(ta, t)
	}

	return ta, err
}
