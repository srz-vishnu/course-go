package main

import "fmt"

func main() {
	var colors map[string]string //one method of declaring map

	//colors := make(map[string]string) ////2nd method of declaring map,15 colors["state"] = "washington"

	// colors := map[string]string{
	// 	"name":    "john",
	// 	"country": "USAaaaa",
	// }

	colors["state"] = "washington"
	//delete(colors, "state")

	//fmt.Println(colors)
	insideMap(colors)
}

func insideMap(c map[string]string) {
	for key, val := range c {
		fmt.Println(key, " is ", val)
		////fmt.Println(c["name"], " is ", c["country"])   //without for we can print like this
	}
}
