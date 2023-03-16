package main

import (
	"fmt"
)

type item struct {
	no    int
	name  string
	price float64
}

func main() {

	var item1 item

	item1.no = 1
	item1.name = "laptop"
	item1.price = 75000.121

	fmt.Println("the item no is:", item1.no)
	fmt.Println("the item name is:", item1.name)
	fmt.Println("the item price is:", item1.price)

}
