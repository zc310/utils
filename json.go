package utils
import (
	"encoding/json"
	"fmt"
)

func PrintJSON(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
func PrintJSONIndent(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
