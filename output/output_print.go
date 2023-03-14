package main

import (
	"fmt"
)

// ques1: a=go b=lang, output =golang venam with no newline and space
// ques2: a=go b=lang, output =golang venam with newline and space
// ques3: a=go b=lang, output ="go" in one line next line will be "lang" venam
// ques4: a=go b=lang, output ="go" in one line next line will be "lang" but no space

func main() {

	var a = "go"

	var b = "lang"

	//fmt.Print(a, b)                // ques1 answer
	fmt.Print(a, b, "\n")       // ques2 answer
	fmt.Print(a, "\n", b, "\n") //ques3 answer
	fmt.Print(a, "\n", b)       //ques4 answer

}
