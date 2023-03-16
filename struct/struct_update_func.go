package main

import "fmt"

type door struct {
	position string
}

func main() {

	home := door{
		position: "open",
	}

	home.updatedoor("closed")
	//fmt.Println("door is:", home.position)   instead of println am using func to print op, so we created printt() func
	home.printt()
}

func (dook *door) updatedoor(newposition string) {
	(*dook.position).position = newposition
}

func (d door) printt() {
	fmt.Println("door is:", d)
}

// func (d door) move() {
// 	// Open door
// 	d.position = "open"
// 	fmt.Printf("Door position is changed to %s\n", d.position)
// }

// func main() {
// 	door1 := door{"closed"}
// 	door1.move()
// 	fmt.Printf("Door is %s\n", door1.position)
// 	//Expected: "open", Result: "closed
// }
