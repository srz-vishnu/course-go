package main

import (
	"fmt"
)

type Person struct { //person struct with 2 fields
	firstname  string
	secondname string
}

type Details struct { // Details struct with 2 fields, actually we are using name as 3rd field and name has 1st n 2nd name
	id     int
	course string
	name   Person
}

func main() {
	User5 := Details{
		id:     2,
		course: "Golang",
		name: Person{
			firstname:  "john",
			secondname: "wick",
		},
	}
	fmt.Println("information is:", User5)
	fmt.Println("id is:", User5.id)
	fmt.Println("course is:", User5.course)
	fmt.Println("1st name is:", User5.name.firstname)
	fmt.Println("2nd name is:", User5.name.secondname)

}
