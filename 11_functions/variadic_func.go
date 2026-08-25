package main

import "fmt"

func sumOfNums(val ...int) int {
	sum := 0
	for _, i := range val {
		sum += i
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumOfNums(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(sumOfNums(nums...))
}

/*
But there's an important rule: the variadic parameter must be the last parameter.

This is valid:
func test(name string, nums ...int) { ✅
}

This isn't:
func test(nums ...int, name string) { // ❌
}
*/
