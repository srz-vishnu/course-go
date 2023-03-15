package main

import (
	"fmt"
)

//ques create nil and empty map

func main() {

	var a = make(map[string]string)

	var b = map[string]string{}

	var c map[string]string // this is nil

	fmt.Println("a: ", a)

	fmt.Println(a != nil)

	fmt.Println("b: ", b)

	fmt.Println(b != nil)

	fmt.Println("c: ", c)

	fmt.Println(c != nil)

}
