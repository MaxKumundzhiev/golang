package main

import "fmt"

/*

*/


func main() {
	i, j := 42, 2701
	fmt.Println(i, j)
	// if u want to get the address of a variable, u can use the & operator
	fmt.Println(&i, &j)

	// we can store the address of a variable in a dedicated variable using same &
	iPointter := &i
	jPointter := &j
	fmt.Println(iPointter, jPointter)

	// with the pointers we can use * operator as well but in 2 different ways
	
	// 1. если стоит перед переменной типа указатель то это оператор разыменования (*somePointer)
	// расшифровывает значение лежащее по адресу (в указателе) и возвращает его
	// also it is called dereferencing
	fmt.Println(*iPointter, *jPointter)

	// we can also change the value of the variable by dereferencing the pointer
	*iPointter = 100
	fmt.Println(i)
	fmt.Println(*iPointter)

	// we can also change the value of the variable by dereferencing the pointer
	*jPointter = 200
	fmt.Println(j)
	fmt.Println(*jPointter)

	// 2. если стоит перед типом данных то это оператор декларации (*somePointerType)
	// создает переменную типа указатель и присваивает ей адрес переменной
	var pointerInMemory *int = new(int)
	fmt.Println(pointerInMemory)
	fmt.Println(*pointerInMemory)

	value := 500
	squareValue(&value)
}



func squareValue(val *int) {
	*val *= *val
	fmt.Println(val, *val)
}