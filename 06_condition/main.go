package main

import "fmt"

func main() {
	var i int = 1
	if i < 18 {
		fmt.Println("you are not 18")
	} else {
		fmt.Println("you adult")
	}

	if true || true {
		fmt.Println("true")
	} else if true && false {
		fmt.Println("false")
	}
}
