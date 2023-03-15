package main

import (
	"fmt"
)

// add fourth as key and value as johnny
// replace second value to roy

func main() {
	y := map[string]string{"first": "jake", "second": "john", "third": "james"}
	fmt.Println("map elements:", y)
	y["fourth"] = "johnny"
	fmt.Println("map elements:", y)
	y["second"] = "Roy"
	fmt.Println("map elements:", y)

}
