package main

import (
	"fmt"
)

func main() {
	hello("vk")
	hello("ajith")
}

func hello(name string) {
	fmt.Println(name, "says hey hello")
}
