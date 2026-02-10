package main

import (
	"fmt"
	"strings"
)


func Printer(str string) {
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Println(string(str[i]))
	}
}

func concatBuider(str string) string {
	// string builder
	var builder strings.Builder
}


func main() {
	Printer("abc")
}
