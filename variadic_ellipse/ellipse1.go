package main

import "fmt"

func daysInWeek(days ...string) { // ellipse function

	for i, d := range days {

		// fmt.Printf("Days %d: %s\n", i+1, d)
		fmt.Println("index:", i, "days", d)

	}

}

func main() {

	fmt.Println("==============================")

	daysInWeek()

	fmt.Println("==============================")

	daysInWeek("Monday")

	fmt.Println("==============================")

	daysInWeek("Monday", "Tuesday")

	fmt.Println("==============================")

	daysInWeek("Monday", "Tuesday", "Wedsday")

}
