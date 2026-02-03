package main

import (
	"fmt"
	"strconv"
)


func singleArgument(argument int) int {
	fmt.Println("single argument function declaration")
	return argument
}


func multipleArgument(argumenntA int, argumentB string) string {
	fmt.Println("multiple argument function declaration")
	concatenated := strconv.Itoa(argumenntA) + argumentB
	return concatenated
}

func namedReturn(request string) (response string) {
	// response is reserved within function scope and automatically
	// returns the reserved variable if not other was explicitly returned
	fmt.Println("single argument function declaration with named return")
	return
}

func multipleVariablesReturn(statusCode int) (string, error) {
    if statusCode == 200 {
        return "success", nil
    }
    return "failed", fmt.Errorf("status code: %d", statusCode)
}

func sliceSum(slice []int) (sum int) {
	for _, num := range slice {
		sum += num
	}
	return 
}


func main() {
	resA := singleArgument(10)
	fmt.Println(resA)

	resB := multipleArgument(10, "20")
	fmt.Println(resB)

	resC := namedReturn("hello this is request")
	fmt.Println(resC)

	fmt.Println(multipleVariablesReturn(200))
	fmt.Println(multipleVariablesReturn(400))

	// var slice []int = []int{1, 2, 3, 4}
	slice := []int{1, 2, 3, 4}
	fmt.Println(sliceSum(slice))

	/* 
	anonymys function
	might be defined without name
	func (input: string) {
		fmt.Println("im called and im anonymys")
	} ("input to anon function")

	or u can define to variable
	printer := func (input: string) {
		fmt.Println("im called and im anonymys")
	}
	
	printer("input to anon function")
	*/


}
