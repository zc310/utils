package utils

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWideString_MarshalJSON(t *testing.T) {
	var a struct {
		A WideString `json:"a"`
	}
	err := json.Unmarshal([]byte(`{"a":"\u6c49\u5b57"}`), &a)
	assert.Equal(t, nil, err)
	assert.Equal(t, a.A, WideString("汉字"))
	b, err := json.Marshal(a)
	assert.Equal(t, nil, err)
	assert.Equal(t, string(b), `{"a":"\u6c49\u5b57"}`)
}
func TestCommaTextToDynInt(t *testing.T) {
	a := CommaTextToDynInt("1,2,3,,4,5,6,7")
	assert.Equal(t, len(a), 7)
	assert.Equal(t, a[6], 7)

	assert.Equal(t, DynIntToCommaText(a), "1,2,3,4,5,6,7")
}
