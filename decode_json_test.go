package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Raden","MiddleName":"Ario","LastName":"Damar","Age":30,"Married":true}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}

	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
}
