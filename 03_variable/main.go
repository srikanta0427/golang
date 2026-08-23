package main

import "fmt"

func main() {
	var name string = "golang"
	name = "golang"
	fmt.Println(name)

	// infer
	var name1 = "golang"
	var isCorrect = true
	fmt.Println(name1,isCorrect)

	var age int = 12
	fmt.Println(age)

	// shorthand syntax
	value:=12
	fmt.Println(value)

	var price float32
	price = 12.5
	fmt.Println(price)
}
