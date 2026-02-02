package main

import (
	"fmt"
)

func main() {
	// bool
	var exists bool = true

	// string
	var name string = "Bob"

	// int (int8  int16  int32  int64 uint uint8 uint16 uint32 uint64 uintptr)
	var number int8 = 100

	// float32 float64
	var float_number float32 = 3.14

	// byte (alias for uint8)
	var b1 byte = 97       // Numeric value (ASCII 'a')
    b2 := byte('b')        // Character literal
    b3 := 'c'              // Default type is rune (int32), but can be converted to byte if within rang

	// rune (alias for int32)
	var r1 rune = 'A' // rune alias for int32 - is meant to work with non ASCII chars

	// complex64 complex128
	var c1 complex64 = complex(3, 2) // 3 + 2i

	fmt.Println(exists, name, number, float_number, b1, b2, b3, r1, c1)
}
