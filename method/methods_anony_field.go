package main

import "fmt"

type parent struct {
	firstName string

	lastName string
}

type student struct {
	name string

	id int

	class int

	parent
}

func (p parent) printParent() {

	fmt.Printf("%s %s is the name of a parent \n", p.firstName, p.lastName)

}

func main() {

	st1 := student{

		name: "Rowland Dave",

		id: 20,

		class: 4,

		parent: parent{

			firstName: "Andrew",

			lastName: "Greg",
		},
	}

	st1.printParent()

	//st1.parent.printParent()

	// fmt.Println(st1.parent) method ila verum normal struct anenkil igane print akam

}
