package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {

	customer := Customer{
		FirstName:  "Raden",
		MiddleName: "Ario",
		LastName:   "Damar",
		Age:        30,
		Married:    true,
		Hobbies:    []string{"Billiard", "Coding", "Gaming"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Raden","MiddleName":"Ario","LastName":"Damar","Age":30,"Married":true,"Hobbies":["Billiard","Coding","Gaming"]}`
	jsonBytes := []byte(jsonString)

	customer := Customer{}
	json.Unmarshal(jsonBytes, &customer)
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}
