package main

import "fmt"

func main () {
	printMessage := true

	if printMessage {
		fmt.Println("Message printed!")
	}

	accountBalance := 20
	accountCredit := 100

	if accountBalance + accountCredit > 0 {
		fmt.Println("There is money available.")
	} else {
		fmt.Println("You don't have any money left.")
	}

	testScore := 50

	if testScore >= 100 {
		fmt.Println("Congrats! You have the top mark!")
	} else if testScore >= 90 {
		fmt.Println("High mark with distinction!")
	} else if testScore >= 70 {
		fmt.Println("High medium mark")
	} else if testScore >= 60 {
		fmt.Println("Medium mark")
	} else if testScore >= 50 {
		fmt.Println("Passed")
	} else if testScore >= 0 {
		fmt.Println("You did not pass!")
	} else {
		fmt.Println("Your score is not correct.")
	}

	hasMoney := true
	hasCredit := true

	if hasMoney && hasCredit {
		fmt.Println("You have money and credit!")
	}

	hasApple := true
	hasBanana := true

	if hasApple || hasBanana {
		fmt.Println("You have some fruit.")
	}

	hasPineapple := false

	if !hasPineapple {
		fmt.Println("You don't have pineapple.")
	}

}
