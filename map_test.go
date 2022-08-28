package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P000001","name":"Samsung Note 10 Plus",
	"price":13000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    "P000001",
		"name":  "Samsung Note 10 Plus",
		"price": 13000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
