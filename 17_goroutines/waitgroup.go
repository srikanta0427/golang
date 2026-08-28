package main

import (
	"fmt"
	"sync"
)

func tasks[T int | string](arg T, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("tasks started", arg)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go tasks("Hello", &wg)
	go tasks("world", &wg)
	wg.Wait()
	fmt.Println("tasks finished")
}
