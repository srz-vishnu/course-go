package main

import (
	"fmt"
)

type Plant interface {
	Grow() string
}

type Mango struct {
}

func (m Mango) Grow() string {
	return "mango can grow"
}

type Apple struct {
}

func (a Apple) Grow() string {
	return "apple can grow"

}

type Orange struct {
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

//vfbb
