package main

import (
	"fmt"
)

//ques crete a  slice from array of value 10,20,30,40,50,60,70 with index 1 to 4 and find length n capacity

func main() {

	numarr := [7]int{10, 20, 30, 40, 50, 60, 70} //array
	//fmt.Printf("Array = %v\n", numarr)
	fmt.Println("array elements are:", numarr)

	num := numarr[1:4]
	fmt.Println("slice elements are:", num)
	fmt.Println(len(num))
	fmt.Println(cap(num))

}
