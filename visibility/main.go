package main

import (
	"fmt"
	"visibility/src/person"
)

func main () {
	person := person.NewPerson(10, "Bob")
	fmt.Println(person)

}