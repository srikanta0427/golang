package main

import (
	"fmt"
)

type Number interface {
	int | float64
}

func Sum[T Number](m T) {
	fmt.Println(m)
}

func Add[T int | float64](a, b T) T {
	return a + b
}

func Prints[T any](a []T) {
	fmt.Println(a)
}

func main() {
	a := []int{5, 6, 4}
	fmt.Println(Add(1, 2))
	fmt.Println(Add(11.2, 233.88))

	Prints([]int{6, 7, 5})
	Prints(a)
	Sum(11)

}
