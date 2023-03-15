package main

import "fmt"

// ques using continue print odd num bte 5 to 15

func main() {

	for i := 5; i <= 15; i++ {

		if i%2 != 1 {

			continue

		}

		fmt.Println(i)

	}

}
