package main

import (
	"fmt"
)

func task1() string {
	return "task 1 completed"
}

func task2() string {
	return "task 2 complete"
}

func worker(ch chan string) {
	ch <- task1()
	ch <- task2()
	ch <- "task 3 completed"
}

func main() {
	var ch = make(chan string, 2)

	go worker(ch)
	//msg := <-ch
	fmt.Println("before work")
	fmt.Println(<-ch)
	//fmt.Println(msg)
	fmt.Println("after work")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
