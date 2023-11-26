package main

import (
	"fmt"
)

func main () {
	// arrays
	// Create an array of capacity 5 and add two elements
	a := [5]int{1,2,3,4,5} // [1,2,0,0,0]
	fmt.Println(a)

	// To declare:
	var b [5]int // [0,0,0,0,0]
	fmt.Println(b)

	var c [5]string // [     ]
	fmt.Println(c)

	// An array's lenght is part of its type, so it cannot be resized
	fmt.Println("Array's length", len(c))
	fmt.Println("Array's capacity", cap(c))

	// Slices
	a_slice := a[1:3]
	fmt.Println(a_slice)
	fmt.Println(a[3])
	fmt.Println("Slice's length", len(a_slice))
	fmt.Println("Slice's capacity", cap(a_slice))

	// Modify an element of the slice
	a_slice[0] = 99

	// As the slice is just a view of the array
	// the modified element is propagated to the
	// original array
	fmt.Println(a_slice)
	fmt.Println(a)

	// Array literal
	arr_bools := [3]bool{true, true, false}
	fmt.Println(arr_bools)
	// Slice literal
	// This creates an array and then builds an slice that
	// references it
	arr_bools2 := []bool{true, true, false}
	fmt.Println(arr_bools2)

	// NIL slice
	// A nil slice has a length and capacity of 0
	// and has no underlying array
	var nil_slice []int
	// Cannot do:
	// nil_slice := []int

	fmt.Println(nil_slice, len(nil_slice), cap(nil_slice))
	if nil_slice == nil {
		fmt.Println("NIL slice")
	}


	// Creating a Slice with make
	// func make(t Type, size ...IntegerType) Type
	//
	// Slice with len 5 and capacity 5
	my_slice := make([]int, 5)
	fmt.Println(my_slice, len(my_slice), cap(my_slice))

	// Slice with len 0 and capacity 0
	my_slice2 := make([]int, 0)
	fmt.Println(my_slice2, len(my_slice2), cap(my_slice2))

	// Slice with len 2 and capacity 5
	my_slice3 := make([]int, 2, 5)
	fmt.Println(my_slice3, len(my_slice3), cap(my_slice3))

	// Slices can contain any type, even other slices
	matrix := [][]int{
		[]int{1, 0},
		[]int{0, 1},
	}
	fmt.Println(matrix)

	// matrix using make
	matrix2 := make([][]int, 2, 2)
	fmt.Println(matrix2)

	// Appending to a slice
	// func append(s []T, vs ...T) []T
	my_slice3 = append(my_slice3, 3)
	fmt.Println(my_slice3)

	// Looping a slice
	for i, v := range my_slice3 {
		fmt.Println(i, v)
	}

	// Range usage alternatives
	// for i, _ := range my_slice3
	// for _, v := range my_slice3
	// for i := range my_slice3 (get only the index)
}
