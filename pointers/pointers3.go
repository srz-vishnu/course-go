package main

import "fmt"

func main() {

	var a = 23

	var p = &a

	p1 := *p + 5

	fmt.Println(p1)

}
