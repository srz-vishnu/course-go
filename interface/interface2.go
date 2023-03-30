package main

import (
	"fmt"
)

type Plant interface { // step1 interface created
	Grow() string // mentioning the method here step 4
}

type Mango struct{} //step 2 is to create a struct
type Orange struct{}
type Apple struct{}

func (m Mango) Grow() string { //step 3 method is created for each struct
	return "mango can grow"
}

func (a Apple) Grow() string {
	return "apple can grow"

}

func (o Orange) Grow() string {
	return "orange can grow"

}

func main() {
	plants := []Plant{Mango{}, Orange{}, Apple{}}
	for _, v := range plants {
		fmt.Println(v.Grow())

	}

}
