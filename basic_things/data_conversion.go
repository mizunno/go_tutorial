package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main () {
	// Print type of a variable
	fmt.Printf("%T\n", os.Args[0])
	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("%T\n", 1)
	fmt.Printf("%T\n", -1)
	fmt.Printf("%T\n", 3.14)

	// It does NOT work using fmt.Println
	// This will print "%T 1"
	fmt.Println("%T", 1)

	// You could try to coerce the type it will raise an error
	// var num int = "1"

	// Another to find the type is by using the reflect package
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(os.Args))

	// Parse string to int
	// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	num, err := strconv.Atoi(os.Args[1])
	fmt.Println("Value returned by strconv: ", num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}

	// Another way
	num2, err2 := strconv.ParseInt("23", 10, 32)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("num2: ", num2)
	}

	// Integer to String
	var numOfCards = 60
	numOfCardsStr := strconv.Itoa(numOfCards)

	fmt.Println("numOfCards as str", numOfCardsStr)
}
