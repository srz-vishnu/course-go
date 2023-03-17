package main

import (
	"fmt"
)

type Employee struct {
	Name  string
	Email string
}

func (e Employee) GetEmail() string {

	e.Email = "newemail@gmail.com"
	return e.Email
}

func (e *Employee) GetName() string {

	e.Name = "vishnu k rajan"
	return e.Name
}

func main() {

	details := Employee{
		Name:  "vishnu",
		Email: "vishnuk@smartrabbitz.com",
	}

	fmt.Println("The value that we have is Name:", details.Name, "and Email:", details.Email)

	details.GetEmail()
	fmt.Println("The value that we have after func call is, Name:", details.Name, "and Email:", details.Email)

	details.GetName()
	fmt.Println("The value that we have after func call is, Name:", details.Name, "and Email:", details.Email)

}
