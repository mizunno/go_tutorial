package main

import (
	"fmt"
)

func main() {
	// Main function serves as entrypoint to the program
	fmt.Println("Main function")
}

func add(num1 int, num2 int) int {
	// Function add receive two ints and return one int
	return num1 + num2
}

// Named return
func product(num1 int, num2 int) (product int) {
	product = num1 * num2
	return
}

// Multiple returns
func calc(num1 int, num2 int) (sum int, product int) {
	sum = num1 + num2
	product = num1 * num2
	return
}
