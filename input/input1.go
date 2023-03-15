package main

import "fmt"

func main() {

	var country string

	var id int

	fmt.Println("Enter Your country name and id: ")

	fmt.Scanln(&country, &id)

	fmt.Printf("country is %s and id is %d\n", country, id)

}
