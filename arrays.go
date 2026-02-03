package main

import "fmt"

func main() {
	// arrays
	var ArrayA [3]int // [0, 0, 0]

	const size = 10
	/*
		can not be used VAR size while inialisation
		due to mutability of "var" size
	*/
	var ArrayB [size]int // [0, 0, ..., 0]

	ArrayC := []int {1, 2, 3}

	fmt.Println(ArrayA)
	fmt.Println(ArrayB)
	fmt.Println(ArrayC)

	// slices (has lenght and capacity properties)
	var SliceA []int  // size=0, capacity=0	
	fmt.Println(len(SliceA), cap(SliceA))

	SliceB := make([]int, 5, 10)  // size=0, capacity=0
	fmt.Println(SliceB)
	fmt.Println(len(SliceB), cap(SliceB))

	SliceB = append(SliceB, )
	
}

