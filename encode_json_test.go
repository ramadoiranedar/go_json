package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func logJson(data interface{}) {
	bytes, err := json.Marshal(data) // convert to json
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes)) // bytes to string for json
}

func TestEncode(t *testing.T) {
	logJson("Damar")
	logJson(6)
	logJson(true)
	logJson([]string{"raden", "ario", "damar"})
}
