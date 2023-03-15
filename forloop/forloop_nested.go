package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {

		fmt.Println(i)

		for j := 1; j <= 2; j++ {

			fmt.Println("I am Inner Loop") // it excutes inner first after it completeing it goes to outer loop excecution

		}

	}

}
