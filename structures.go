package main

import "fmt"

/*
how to define a structure

type <name> struct {
	<field1> <type1>
	<field2> <type2>
	...
}

structure is a collection of fields, each field has a name and a type
structure is a user-defined type that can be used to group related data together
structure can be used to create complex data types that can represent real-world entities
*/


type User struct {
	Id int
	Name string
	Age int
}

type Account struct {
	Id int
	Name string
	Owner User
}


func main () {
	var user User = User{
		Id: 1,
		Name: "Max",
		Age: 29,
	}
	
	fmt.Println(user)

	var account Account = Account{
		Id: 1,
		Name: "Max's account",
		Owner: user,
	}
	fmt.Println(account)

}
