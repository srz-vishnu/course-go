package main

import (
	"fmt"
)

type Student struct {
	name string

	id int

	class int
}

type Parent struct {
	firstName string

	lastName string

	childID int
}

func (s Student) printDetails() {

	fmt.Printf("%s with ID number %d is in class %d \n", s.name, s.id, s.class)

}

func (p Parent) printDetails() { //same method as above

	fmt.Printf("%s %s is the parent of a student with the ID number %d \n", p.firstName, p.lastName, p.childID)

}

func main() {

	st1 := Student{

		name: "Rowland Dave",

		id: 20,

		class: 4,
	}

	st1.printDetails()

	pt1 := Parent{

		firstName: "Dave",

		lastName: "James",

		childID: 20,
	}

	pt1.printDetails() //metods with same name but different instance ie instance of parent and student

}
