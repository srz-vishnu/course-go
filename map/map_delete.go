package main

import (
	"fmt"
)

// ques delete the 2nd element

func main() {
	x := map[int]string{1: "vk", 2: "ajith", 3: "suji"}
	fmt.Println("elements of map is:", x)
	delete(x, 3)
	fmt.Println("elements of map is:", x)

}
