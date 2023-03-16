package main

import (
	"fmt"
)

type User struct {
	string // here we dont specified any field name, just mentioned datatype...this is anonymous field struct
	int
}

func main() {

	user1 := User{

		string: "Kate", // here also value assigning is different, since no fieldname datatype name only used
		int:    65,
	}

	fmt.Println(user1.string)
	fmt.Println(user1.int)
	fmt.Println("detaisl is:", user1)

}
