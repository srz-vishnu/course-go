package main

import (
	"fmt"
)

type User struct {
	name string
	id   int
	fee  int
}

func main() {
	user5 := User{
		name: "james",
		id:   3,
		fee:  500,
	}
	fmt.Println("user5 details before is:", user5)

	// changing user details
	user5.name = "roy"
	user5.fee = 600
	fmt.Println("user5 details after is:", user5)

}
