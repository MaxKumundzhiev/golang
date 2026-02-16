package main

/*
в го мы работаем с обьектами которые лежат в памяти в ячейке.
у каждой ячейки есть свой адрес.

в функции можно передавать переменные по значению или по указателю (ссылке)


// по значению
func increase(value int) {
	value++
}

value := 10
increase(value)

// по указателю
func increase(value *int) {
	*value++
}
increase(&value)

*/

import "fmt"
 
func increase(value int) {
	/*
	В Go по умолчанию значения передаются в функции по копии. 
	Это значит, что при передаче аргумента в функцию будет создаваться копия объекта. 
	Поэтому изменения внутри функции не будут видны снаружи (потому что была изменена локальная копия).
	*/
	value++
}


type MutableStruct struct {
	value int
}

func (s *MutableStruct) SetValue(newValue int) {
	s.value = newValue
}


type ImmutableStruct struct {
	value int
}

func (s ImmutableStruct) SetValue(newValue int) {
	s.value = newValue
}



func main() {
	value := 5
	fmt.Printf("before %d\n", value)
	increase(value)
	fmt.Printf("after %d\n", value)

	mutable := MutableStruct{10}
	immutable := ImmutableStruct{10}

	mutable.SetValue(11)
	immutable.SetValue(11)
	fmt.Println(mutable) // 11
    fmt.Println(immutable) // 10
}