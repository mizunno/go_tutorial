package main

import (
	"fmt"
)

func main () {
	// Read user input
	var user_input string
	fmt.Println("Enter something: ")
	fmt.Scan(&user_input)
	fmt.Println(user_input)

	// Read multiple inputs
	var input1, input2, input3 string
	fmt.Println("Enter 3 things: ")
	// You can input the values separated by white spaces
	// or by newline
	fmt.Scan(&input1, &input2, &input3)
	fmt.Println("input1: ", input1)
	fmt.Println("input2: ", input2)
	fmt.Println("input3: ", input3)

	// Read user input with formatting
	var prefix string
	var num int
	fmt.Println("Enter the invoice id: ")
	fmt.Scanf("%3s%d", &prefix, &num)
	fmt.Printf("Invoice prefix: %s \nInvoice number: %d", prefix, num)
}
