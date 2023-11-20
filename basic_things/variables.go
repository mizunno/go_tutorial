package main

import "fmt"

// Declare variables. Explicit way of giving a type
// to a variable
var variableName string

var (
	firstName = "Chris"
	age = 20
)


func main () {
	// Assign value to an existing variable. Only within functions.
	variableName = "variableValue"

	// Define and assign a value. Only within functions
	// Type is inferred
	name := "Peter"
	fmt.Println(name)

	var (
		customerName = "Nicole"
		accountBalance = 20
	)

	// String interpolation
	fmt.Printf("Customer %s has %d euros on their bank account", customerName, accountBalance)
}
