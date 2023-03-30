package main

import "fmt"

// Author structure
type author struct {
	name      string
	branch    string
	particles int
}

// Method with a receiver of author type
func (a *author) show(print string) {
	(*a).branch = print
}

func main() {

	// assigining val to struct
	res := author{
		name:   "Sona",
		branch: "CSE",
	}

	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch Name(Before): ", res.branch)

	// pointer
	p := &res

	p.show("ECE") //assigining new branch name
	fmt.Println("Author's name: ", res.name)
	fmt.Println("Branch Name(After): ", res.branch)
}
