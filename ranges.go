package main

import "fmt"

/*
range iterates over elements in a variety of data structures.


Unknown type → use %v
Strings → %s
Int → %d
Float → %f
Debug structs → %+v
Print type → %T
*/


func TraversalSumOverArrayLike() {
	// Here we use range to sum the numbers in a slice.
	/*
	range on arrays and slices provides both the index and value for each entry. 
	Above we didn’t need the index, so we ignored it with the blank identifier _. 
	Sometimes we actually want the indexes though.
	*/

	slice := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, num := range slice {
		sum += num
	}
	fmt.Println("sum:", sum)
}

func TraversalLookUpOverMap() {
	// range on map iterates over key/value pairs.

	hashmap := map[string]string{"1": "Bob", "2": "Foo"}
	for key, value := range hashmap {
		if value == "Foo" {
			fmt.Printf("Found user %s with id %s\n", value, key)
			return
		}
	}
	fmt.Println("User not found")
}

func TraverseOverString() {
	// range on strings iterates over Unicode code points. 
	// The first value is the starting byte index of the rune and the second the rune itself.

	sentence := "Hellow друг!"
	for idx, rune := range sentence {
		fmt.Printf("Idx %d  unicode %c\n", idx, rune)
	}
}


func main() {
	TraversalSumOverArrayLike()
	TraversalLookUpOverMap()
	TraverseOverString()
}
