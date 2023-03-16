package main

import "fmt"

type washingMachine interface {
	wash()

	dry()
}

type samsung struct {
	kg int
}

func (w samsung) wash() {

	fmt.Println("This machine can WASH")

}

func (w samsung) dry() {

	fmt.Println("This machine can DRY")

}

func main() {

	var a washingMachine

	a = samsung{kg: 15} // instance of an struct final step

	a.wash()
	a.dry()

}

// first we need to create a interface
// then a struct
// we need to connect struct to interface, for that we are using methods
// creating 2 methods
