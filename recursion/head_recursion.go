package main

import (
	"fmt"
)

func count(n int) {

	if n > 0 {

		count(n - 1) /*head recursive, because this is the fist statement in the function, if it is under print as last stmt then its called tail*/

		fmt.Println(n)

	}

}

func main() {

	count(10)

}
