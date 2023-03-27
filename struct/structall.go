package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//alex := person{"sam", "antonnny", 25}  //one method
	//fmt.Println("detail is :", alex)

	// alex := person{firstName: "sam", lastName: "antony", age: 25}  //2nd method
	// fmt.Println("detail is :", alex)

	// var alex person // variable alex is of type person  // 3rd method
	// alex.firstName = "sam"
	// alex.lastName = "antony"
	// alex.age = 25
	// fmt.Println("details is:", alex)
	// fmt.Printf("detail is: %+v", alex) //%+v print alex with all values n assigneee,next 4th method with one more struct combaining

	alex := person{
		firstName: "sam",
		lastName:  "antony",
		age:       25,
		contact: contactInfo{
			email:   "sam@gmail.com",
			zipCode: 589632,
		},
	}

	//fmt.Printf("%+v", alex) igane ane if there is no func under
	// alex.printt()
	// alex.updateMe("jack") //update cheyan noki but update ayilla,have to crct it using pointer

	alexPointer := &alex
	alexPointer.updateMe("jack")
	alex.printt()
}

func (p person) printt() {
	fmt.Printf("%+v", p)
}

// func (p person) updateMe(newfirstName string) {  //ivide pointer ila so not updating
// 	p.firstName = newfirstName   // so func vendum ezhuthunu with pointer
// }

func (pointerToPerson *person) updateMe(newfirstName string) {
	(*pointerToPerson).firstName = newfirstName
}
