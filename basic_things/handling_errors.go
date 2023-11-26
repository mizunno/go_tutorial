package main

import (
	"fmt"
	"errors"
)

func main() {
	var nominator int
	var divider int

	// Using panic / recover
	fmt.Println("Enter nominator and divider:")
	fmt.Scan(&nominator, &divider)

	result := Divide(nominator, divider)

	fmt.Println("Divide result:", result)

	fmt.Println("-------------------------")
	// Using error pattern
	fmt.Println("Enter nominator and divider:")
	fmt.Scan(&nominator, &divider)

	result, error := Divide2(nominator, divider)

	if error != nil {
		fmt.Println("Error using divide2:", error)
	} else {
		fmt.Println("Divide2 result:", result)
	}

}

func Divide(nominator int, divider int) float32 {
	defer errorHandler()
	if divider == 0 {
		panic("Divison by 0.")
	}

	// we have to parse int to float32 because
	// int divided by another int produces an int
	// https://go.dev/ref/spec#Constant_expressions
	return float32(nominator) / float32(divider)
}

func errorHandler() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

// Error pattern with errors package
func Divide2(nominator int, divider int) (float32, error){
	if divider == 0 {
		return 0, errors.New("divider cannot be 0")
	}

	return float32(nominator) / float32(divider), nil
}
