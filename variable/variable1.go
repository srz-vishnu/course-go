package main

import (
	"fmt"
	"reflect"
)

func main() {

	var name string = "john" //differnt type of variable declaration
	var name1 = "roy"
	name2 := "alex"
	var price int
	var country string

	price = 50
	country = "UK"

	fmt.Println("the name is:", name)
	fmt.Println("the name is:", name1)
	fmt.Println("the name is:", name2)
	fmt.Println("the price is:", price)
	fmt.Println("the country is:", country)

	fmt.Println(reflect.TypeOf(name1))
	fmt.Println(reflect.TypeOf(price))

}
