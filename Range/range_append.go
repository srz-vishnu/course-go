package main

import "fmt"

// we have lang slce, we have to append daniel 4 times using range

func main() {

	lang := []string{"Java", "Python", "Golang", "C++"}

	for range lang {

		lang = append(lang, "Daniel")

	}

	fmt.Println("lang before", lang)

	fmt.Println("--------------------------")

	fmt.Printf("%q\n", lang)
}
