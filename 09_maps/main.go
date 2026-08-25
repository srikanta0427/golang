package main

import (
	"fmt"
	"maps"
)

func main() {

	// creating map
	m := make(map[string]int)
	// key value pair
	m["srikanta"] = 22
	m["srikant"] = 21

	// get an element
	fmt.Println(m["srikanta"])
	fmt.Println(len(m))

	// to delete an element in map
	fmt.Println(m)
	delete(m, "srikant")
	fmt.Println(m)

	// to clear map
	clear(m)
	fmt.Println(m)

	// another way of creating map
	mp := map[string]int{"price": 22, "quantity": 4}
	fmt.Println(mp)

	k, ok := mp["price "]
	fmt.Println(k, ok) // value bool
	if ok {
		fmt.Println(mp["price"])
	} else {
		fmt.Println("not found")
	}

	// check that both map are equal or not
	fmt.Println(maps.Equal(m, mp))
}
