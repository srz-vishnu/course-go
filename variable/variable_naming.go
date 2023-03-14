package main

import (
	"fmt"
)

func main() {

	var myCustomerName = "Eric Gath" //camel case

	var MyCustomerName = "Mike" //pascal case

	var my_customer_name = "Anne" //snake case

	fmt.Println(myCustomerName)

	fmt.Println(MyCustomerName)

	fmt.Println(my_customer_name)

}
