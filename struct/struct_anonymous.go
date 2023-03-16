package main

import (
	"fmt"
)

func main() {
	user := struct { // anonymous struct defines like this
		no      int
		name    string
		country string
	}{ // if we start the new bracket on next line it will be error, thats how anonymous struct works
		no:      1,
		name:    "john",
		country: "UK",
	}
	fmt.Println("user:", user)
}
