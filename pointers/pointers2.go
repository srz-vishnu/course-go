package main

import (
	"fmt"
)

func main() {

	k := new(int)

	*k = 90

	fmt.Println(k)

	fmt.Println(*k)

}
