package main

import (
	"fmt"
)

func main() {
	z := make(map[string]string)

	z["First"] = "Dan"

	z["Second"] = "Isaac"

	z["Third"] = "Jane"

	//syntax     val,key_exist := map_name[key]

	val1, key1 := z["Third"]

	val2, key2 := z["position"] //position not exist so false

	val3, key3 := z["First"]

	fmt.Println(val1, key1)

	fmt.Println(val2, key2)

	fmt.Println(val3, key3)
}
