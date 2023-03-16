package main

import (
	"fmt"
)

type Name struct {
	firstName string
	lastName  string
}

type User struct {
	age int
	fee int
	Name
}

func main() {

	user5 := User{

		Name: Name{
			firstName: "James",
			lastName:  "Gim",
		},
		age: 42,
		fee: 200,
	}
	fmt.Println("First Name:", user5.firstName) //Promoted Field
	fmt.Println("Last Name:", user5.lastName)   //Promoted Field
	fmt.Println("Age:", user5.age)
	fmt.Println("Fee:", user5.fee)

}
