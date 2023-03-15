package main

import (
	"fmt"
)

// ques create a map

func main() {

	var x = map[int]string{1: "India", 2: "UK", 3: "USA"}                             //int for key datatype and string for value datatype
	y := map[string]int{"First": 1, "Second": 2, "third": 3, "Fourth": 4, "Fifth": 5} //string for keytype and int for value type
	z := map[string]string{"name": "vk", "city": "ijk"}                               // both key n value are string

	//map using make function

	a := make(map[int]string)
	fmt.Println("map key value pair before:", a)

	a[1] = "golang"
	a[2] = "react"
	fmt.Println("map key value pair after:", a)

	fmt.Println("map key value pair is:", x)
	fmt.Println("map key value pair is:", y)
	fmt.Println("map key value pair is:", z)

}
