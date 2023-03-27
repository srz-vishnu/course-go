package main

import "fmt"

type door struct {
	position string
}

func main() {

	home := door{
		position: "open",
	}

	// home.updatedoor("closed")
	// fmt.Println("door is:", home.position) // instead of println am using func to print op, so we created printt() func
	// home.printt()

	home.updatedoor("closed but open")
	home.printt()

}

// func (d door) updatedoor(newposition string) {
// 	d.position = newposition
// }

func (dook *door) updatedoor(newposition string) {
	(*dook).position = newposition
}

func (d door) printt() {
	fmt.Println("door is:", d)
}
