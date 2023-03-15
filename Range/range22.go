package main

import "fmt"

func main() {

	days := map[string]string{

		"Day 1": "Monday",

		"Day 2": "Tuesday",

		"Day 3": "Wednesday",

		"Day 4": "Thurdsday",

		"Day 5": "Friday",
	}

	fmt.Println("--------First Method-------------")

	for k, v := range days {

		fmt.Println(k, ":", v)

	}

	fmt.Println("--------Second Method-------------")

	for k := range days {

		fmt.Println(k, ":", days[k])

	}

}
