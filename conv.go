package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// GetString convert interface to string.
func GetString(v interface{}) string {
	switch result := v.(type) {
	case string:
		return result
	case []byte:
		return string(v.([]byte))
	case []string:
		return strings.Join([]string(v.([]string)), "\n")
	default:
		if v != nil {
			return fmt.Sprint(result)
		}
	}
	return ""
}

// GetBool convert interface to bool.
func GetBool(v interface{}) bool {
	switch result := v.(type) {
	case bool:
		return result
	default:
		if d := GetString(v); d != "" {
			value, _ := strconv.ParseBool(d)
			return value
		}
	}
	return false
}

// GetInt convert interface to int.
func GetInt(v interface{}) int {
	switch result := v.(type) {
	case string:
		return StrToInt(strings.TrimSpace(result))
	case int:
		return result
	case int64:
		return int(result)
	default:
		if v != nil {
			return StrToInt(fmt.Sprint(result))
		}
	}
	return 0
}
