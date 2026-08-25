package main

import "fmt"

func main() {
	nums := []int{3, 7, 2}

	// here i is indices
	for i, num := range nums {
		fmt.Println(num, i)
	}
	// iterate map

	mp := map[string]int{"price": 22, "quantity": 4}

	// where k->key and v->value
	for k, v := range mp {
		fmt.Println(k, v)
		//fmt.Println(i, m)
	}

	// use of range in string
	for i, s := range "golang" {
		fmt.Println(i, string(s))
	}
	
}
