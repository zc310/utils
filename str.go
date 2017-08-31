package utils

// IfEmpty return first non-empty string
func IfEmpty(a, b string) string {
	if a == "" {
		return b
	}
	return a
}
