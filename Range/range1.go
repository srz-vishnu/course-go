package main

import (
	"fmt"
)

func main() {
	num := [...]int{2, 4, 6, 8, 10}
	fmt.Println("array ele:", num)
	for i, e := range num {
		//fmt.Printf("Element %v is at index %d \n", e, i)
		fmt.Println("element:", e, "is at index:", i)
	}
}
