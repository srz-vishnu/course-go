package main

import (
	"fmt"
)

func main() {

	pricelist("laptop", 50000)
	pricelist("mobile", 25000)

}

func pricelist(item string, price float64) {
	fmt.Println("the price of the", item, "is", price)
}
