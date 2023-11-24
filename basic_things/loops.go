package main

import (
	"fmt"
)

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		// run this 10 times
		fmt.Print(i)
	}

	// while loop
	c := 0
	for c < 10 {
		fmt.Println(c)
		c++
	}

	// for each
	array := []int{1,2,3,4,5}
	for i,v := range array {
		fmt.Printf("index: %d, item: %d \n", i, v)
	}

	keepLooping := true
	var input string
	for keepLooping {
		fmt.Println("To stop looping type 'quit'")
		fmt.Scan(&input)

		if input == "quit" {
			keepLooping = false
		}
	}


}
