package main

import (
	"fmt"
)

func main() {
	age := 17

	if age < 18 {
		fmt.Println("age is less then 18 not eligable")
	} else {
		fmt.Println("person is eligable")
	}
}
