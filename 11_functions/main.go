package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// multiple returns
func getResult() (string, int, bool, float64) {
	return "hello", 1, true, 4.4
}

// passing func as argument
func calculator(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	sum := add(1, 2)
	fmt.Println(sum)

	//fmt.Println(getResult())
	re1, re2, re3, _ := getResult()
	fmt.Println(re1, re2, re3)

	fn := calculator(10, 20, add)
	fmt.Println(fn)

}
