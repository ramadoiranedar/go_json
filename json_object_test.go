package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Raden",
		MiddleName: "Ario",
		LastName:   "Damar",
		Age:        30,
		Married:    true,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Damar",
		Addresses: []Address{
			{
				Street:     "Jl.Pinggir A",
				Country:    "Indonesia",
				PostalCode: "11111",
			},
			{
				Street:     "Jl.Pinggir b",
				Country:    "Indonesia",
				PostalCode: "22222",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Damar","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jl.Pinggir A","Country":"Indonesia","PostalCode":"11111"},{"Street":"Jl.Pinggir b","Country":"Indonesia","PostalCode":"22222"}]}`

	jsonBytes := []byte(jsonString)

	customer := Customer{}
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jl.Pinggir A","Country":"Indonesia","PostalCode":"11111"},{"Street":"Jl.Pinggir b","Country":"Indonesia","PostalCode":"22222"}]`

	jsonBytes := []byte(jsonString)

	addresses := []Address{}
	err := json.Unmarshal(jsonBytes, &addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
	fmt.Println(addresses[1])
	fmt.Println(addresses[1].Street)
	fmt.Println(addresses[1].Country)
	fmt.Println(addresses[1].PostalCode)
}
