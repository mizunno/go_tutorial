package main

import (
	"fmt"
)

// Struct definition
// type name struct {... fields}

type Address struct {
    city	string
	street	string
	postal	string
}

type Person struct {
    name	string
	age		int
	Address // Equivalent to Address Address
}

// Adding behaviour to a Struct
// func (m type_name) function_name() return_type {
func (a Address) string() string {
	return fmt.Sprintf("City: %s, Street: %s, Postal: %s", a.city, a.street, a.postal)
}

func main () {
	// Create an instance of address struct
	var addr Address
	addr.city = "Sevilla"
	addr.street = "Juan Ramon Jimenez"
	addr.postal = "41012"

	fmt.Println("addr", addr)

	// Another way of create an instance of address struct
	addr2 := Address{"Zaragoza", "Federico Garcia Lorca", "50017"}
	fmt.Println("addr2", addr2)

	peter := Person{
		"Peter",
		30,
		addr,
	}

	fmt.Println(peter)
	fmt.Println(peter.Address.string())

}
