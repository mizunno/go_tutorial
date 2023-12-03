package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

// Working with JSON files

type Product struct {
	Id		int		`json: "id"`
	Name	string	`json: "name"`
}

type Products struct {
	Products []Product	`json: "products"`
}

func main () {
	// Read JSON from string
	json_str := `{"id": 3, "name": "Pineapple"}`
	product_id_3 := Product{}
	// Note we convert to a byte array
	json.Unmarshal([]byte(json_str), &product_id_3)
	fmt.Println(product_id_3)

	// Read JSON from file
	file, err := ioutil.ReadFile("products.json")

	if err != nil {
		fmt.Println("Cannot read products.json file.")
	}

	data := Products{}
	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		fmt.Println("Cannot parse products.json file as JSON.")
	}

	fmt.Println(data)
	fmt.Println(data.Products)

	// Write JSON
	// Add new product to data
	data.Products = append(data.Products, product_id_3)

	fmt.Println(data.Products)

	// json.Marshal returns a byte array
	data_json, err := json.Marshal(&data)

	if err != nil {
		fmt.Println("Cannot parse data to JSON")
	}

	// Convert to string the byte array
	fmt.Println(string(data_json))


}
