package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
strings in go are represented as a slice of bytes (BUT THEY ARE NOT SLICE OF BYTES AS DS, THEY ARE STRINGS DS)
where byte represent a symbol or a part of it, depending on encoding (e.g. UTF, ...)
NOTE: Strings are just slices of bytes without capacity.

difference between array and slice in a nutshell
array - fixed size
	var array [3]int = [3]int{1,2,3}

slice - dynamic size
	var slice []int = []int{1,2,3}


when accessing a particular char in a string u-ll get a byte representation
when getting len of string, be aware ull get the cnt of bytes and not char themselves
for that u need

Notes
	- strings are enclosed with double quotes only
	- string may be empty, but not nil
	- values of string type are immutable
*/


func main() {
	name := "Макс"
	fmt.Println(name)
	fmt.Println(len(name))

	for idx, char := range name {
		fmt.Println(idx, char)
	}

	// operations
	// len of string in bytes
	fmt.Println(len(name)) // == fmt.Println(name) both of them gonna return 8 bytes cause each char takes 2 bytes
	fmt.Println(utf8.RuneCountInString(name))
	// len of string in characters
	// correct way to do so is to use unicode package

	// Add a char to string
	/*
	There are 3 main ways:

	1. Concatenate strings using +=
	- inefficient in loops because strings are immutable
	- every operation allocates a new string

	2. Use strings.Builder
	- most efficient for repeated concatenation
	- minimizes allocations

	3. Use []byte (or []rune) slice
	- append to slice, then convert to string at the end
	- efficient and flexible, especially when working with bytes
	*/

	
	// 1
	str1 := "Hello"
	fmt.Println(str1)
	str1 += " world"
	fmt.Println(str1)

	// 2
	str2 := "builder"
	builder := strings.Builder{}
	builder.WriteString(str2) // returns (int, error), where int represents amt of bytes written to builder
	fmt.Println(builder.String())
	builder.WriteString(" hmm")
	fmt.Println(builder.String())

	// 3
	str3 := []rune{1052, 1072, 1082, 1089}
	fmt.Println(str3)
	str3[0] = 'м'
	fmt.Println(string(str3))

	// remove a char from string
	// look up a char in string

}