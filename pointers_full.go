package main

import "fmt"

/*
каждая переменная имеет ячейку в памяти и адрес к этой ячейке.
иногда к переменной мы можем образаться через другие указатели также (напрмер когда передаем в функцию);

в указателях есть 2 основных оператора: & и *
& - оператор взятия адреса 
* - используется 2мя способами
	1 *variable - deferencing (разименование) то есть получени истенного значения которое лежит по этому адресу 
	2 pointer *int - определение переменной типа указатель на тип (в данном случае на int)
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

	fmt.Println(*initPerson())
}


func squareValue(val *int) {
	*val *= *val
	fmt.Println(val, *val)
}


type Person struct{
	Name string
	Age int
}

func initPerson() *Person {
	m := Person{Name: "Max", Age: 29}
	return &m
}