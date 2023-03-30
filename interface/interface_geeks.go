package main

import (
	"fmt"
)

// defining interface

type tank interface {
	Tankarea() float64
	Tankvolume() float64
}

type measurement struct {
	radius float64
	height float64
}

func (m measurement) Tankarea() float64 {
	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m measurement) Tankvolume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

func main() {
	var t tank
	t = measurement{10, 14}
	fmt.Println("tank area is:", t.Tankarea())
	fmt.Println("tank volume is:", t.Tankvolume())
}
