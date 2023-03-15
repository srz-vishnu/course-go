package main

import "fmt"

func calculate(a int, nums ...int) { // so we need to provide atleast one argument at call
	total := 0
	for _, v := range nums { //v is value
		total += v * a

	}
	fmt.Println(total)
}

func main() {

	calculate(10)
	calculate(10, 3)
	calculate(5, 10, 20)
	calculate(100, 3, 2, 1)

}
