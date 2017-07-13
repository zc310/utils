package utils

import (
	"testing"
	"time"

	. "gopkg.in/go-playground/assert.v1"
)

func Test_Time(t *testing.T) {
	t1, err := TimetampToTime("1352397861001")
	Equal(t, err, nil)
	Equal(t, t1, time.Date(2012, 11, 8, 18, 4, 21, int(time.Millisecond), time.UTC))

	t1, err = TimetampToTime("1352397861")
	Equal(t, err, nil)
	Equal(t, t1, time.Date(2012, 11, 8, 18, 4, 21, 0, time.UTC))

	t1, err = TimetampToTime("1351700038292387000")
	Equal(t, err, nil)
	Equal(t, t1, time.Date(2012, 10, 31, 16, 13, 58, 292387000, time.UTC))

}
