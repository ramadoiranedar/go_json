package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P000001",
		Name:     "Samsung Note 10 Plus",
		ImageURL: "https://thisurl.com/samsungnote10plus.jpg",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P000001","name":"Samsung Note 10 Plus","image_url":"https://thisurl.com/samsungnote10plus.jpg"}`
	jsonBytes := []byte(jsonString)

	product := Product{}

	json.Unmarshal(jsonBytes, &product)
	fmt.Println(product)
}
