package main

import "fmt"

/*
Методы в Go - это функции, которые связаны с определёнными типами данных. Они позволяют определять поведение для структур и других типов данных,
что делает код более организованным и удобным для понимания.
Методы в Go определяются с помощью специального синтаксиса, который включает в себя имя типа,
к которому относится метод, и имя метода.

Синтаксис определения метода:

func (receiver Type) MethodName(parameters) returnType {
	// тело метода
}

receiver - это переменная, которая представляет экземпляр типа, к которому относится метод.

Также можно использовать указатель на тип в качестве receiver, что позволяет и
зменять состояние объекта внутри метода.

Пример:

type Person struct {
	Name string
	Age int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}
В этом примере мы определили структуру Person с полями Name и Age, а также метод Greet, который выводит приветствие с именем и возрастом человека.

Методы в Go могут быть вызваны на экземплярах типов, к которым они относятся. Например:

func main() {
	person := Person{Name: "Alice", Age: 30}
	person.Greet() // Output: Hello, my name is Alice and I am 30 years old.
}


Важно понимать что reciever может быть как значением так и указателем, и это влияет на то,
может ли метод изменять состояние объекта.

Если receiver - это значение, то метод работает с копией объекта, и любые изменения внутри метода не повлияют на оригинальный объект.
Если receiver - это указатель, то метод работает с оригинальным объектом, и изменения внутри метода будут отражаться на нём.

Пример с указателем:
type Person struct {
	Name string
	Age int
}

func (p *Person) SetName(name string) {
	p.Name = name
}

*/

type Person struct {
	Name string
	Age int
}

func (p* Person) UpdateName(name string) {
	p.Name = name
}


func main () {
	var person Person = Person{
		Name: "Max",
		Age: 29,
	}
	fmt.Println(person)

	person.UpdateName("Bob")
	fmt.Println(person)
}
