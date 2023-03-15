package main

import (
	"fmt"
)

type Student struct {
	name string

	id int

	class int
}

func printStudentDetails(s Student) {

	fmt.Printf("%s with ID number %d is in class %d \n", s.name, s.id, s.class)

}

func main() {

	st1 := Student{

		name: "Rowland Dave",

		id: 20,

		class: 4,
	}

	printStudentDetails(st1)

}

// METHOD:

// -------

// package main

// import (

// "fmt"

// )

// type Student struct {

// name  string

// id    int

// class int

// }

// func (s Student) printStudentDetails() {

// fmt.Printf("%s with ID number %d is in class %d \n", s.name, s.id, s.class)

// }

// func main() {

// st1 := Student{

// name:  "Rowland Dave",

// id:    20,

// class: 4,

// }

// st1.printStudentDetails()

// }
