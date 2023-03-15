package main

import "fmt"

func pointerAdd(a *int) {

	fmt.Println(*a + *a)

}

func main() {

	var l int = 4

	pointerAdd(&l)

}

// 2 arguments

// package main

// import "fmt"

// func pointerAdd(a *int, b *int) {

// fmt.Println(*a + *b)

// }

// func main() {

// var l int = 4

// var m int = 10

// pointerAdd(&l, &m)

// }
