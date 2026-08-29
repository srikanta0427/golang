package main

import (
	"fmt"
	"sync"
	"time"
)

func tasks[T int | string](arg T, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("tasks started", arg)
	time.Sleep(time.Second * 2)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go tasks("Hello", &wg)
	go tasks("world", &wg)
	wg.Wait()
	fmt.Println("tasks finished")
}
