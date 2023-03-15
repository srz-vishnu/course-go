package main

import (
	"fmt"
)

func count(n int) {

	if n > 0 {

		fmt.Println(n)

		count(n - 1) /*tail recursive, because this is the last statement*/

	}

}

func main() {

	count(10)

}
