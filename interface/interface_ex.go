package main

import "fmt"

// interface
type Shape interface {
	area() float32
}

// struct to implement interface
type Rectangle struct {
	length, breadth float32
}

// reciever parameter = mathod
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

// we are accessing the interface
// func calculate(s Shape) {
// 	fmt.Println("Area:", s.area())
// }

// main function
func main() {

	var rect Shape
	rect = Rectangle{7, 4}
	fmt.Println("value:", rect.area())
}
