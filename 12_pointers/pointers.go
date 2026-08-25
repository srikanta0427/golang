package main

import "fmt"

func changed(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	var num int = 1
	fmt.Println(num)
	fmt.Println(changed(&num), num)
}
