package utils

import (
	"strconv"
	"time"
)

//TimetampToTime string Timetamp to Time
func TimetampToTime(v string) (time.Time, error) {
	t, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	switch len(v) {
	//java.lang.System.currentTimeMillis()
	case 13:
		return time.Unix(0, t*int64(time.Millisecond)).UTC(), nil
	case 19:
		return time.Unix(0, t).UTC(), nil
	}
	return time.Unix(t, 0).UTC(), nil
}
