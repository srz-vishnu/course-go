package main

import (
	"fmt"
)

//ques is to print output like this
//name is peter
// age is 35
//peter is 35 years old

func main() {

	name := "Peter"

	age := 35

	fmt.Printf("Name is: %s\n", name) //%v is also fine

	fmt.Printf("Age is: %d\n", age)

	fmt.Printf("%s is %d years old\n", name, age)

}
