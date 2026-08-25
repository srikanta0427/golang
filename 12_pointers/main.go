package main

import "fmt"

func main() {
	num := 12
	var num2 *int = &num
	fmt.Println(num, *num2, num2, &num)
	*num2 = 10
	fmt.Println(num, *num2, num2, &num)
	num = 11
	fmt.Println(num, *num2, num2, &num)

	// also
	value := 12
	pointValue := &value
	*pointValue = 44
	fmt.Println(*pointValue, value)
}
