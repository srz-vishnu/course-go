package main

import "fmt"

type parent struct {
	id        int
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

	fmt.Printf("%s %s is the name of a parent with id %d\n", p.firstName, p.lastName, p.id)

}

func main() {

	st1 := student{

		name: "Rowland Dave",

		id: 20,

		class: 4,

		parent: parent{

			firstName: "Andrew",

			lastName: "Greg",
			id:       5,
		},
	}

	st1.printParent()

	//st1.parent.printParent()

	// fmt.Println(st1.parent) //method ila verum normal struct anenkil igane print akam

}
