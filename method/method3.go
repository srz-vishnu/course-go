package main

import (
	"fmt"
)

type Student struct {
	name string

	id int

	class int
}

func (s *Student) changeDetails(id, class int) {

	// changes made inside a method with a point reciver are visible to the caller
	s.id = id

	s.class = class

}

func main() {

	st1 := Student{

		name: "Rowland Dave",

		id: 20,

		class: 4,
	}

	fmt.Printf("Before id: %d and Before class: %d \n", st1.id, st1.class)

	st1.changeDetails(35, 11) //changes made inside a method with a point reciver are visible to the caller

	//(&st1).changeDetails(35, 11)  //can use both 47 n 49

	fmt.Printf("After id: %d and After class: %d \n", st1.id, st1.class)

}
