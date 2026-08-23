package main

import "fmt"

const age = 12

var name = "Alice"

// nm:=12 // not possible
func main() {
	const name string = "golang"
	const isCorrect = true

	const (
		port = 5000
		host = "127.0.0.1"
	)
	fmt.Println(port,host)
}
