package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// classical loop
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	i = 11
	// using range
	for i := range 40 {
		fmt.Println(i)
	}
	
	for i = 1; i <= 100; i++ {
		fmt.Println(i)
	}
}
