package main

import (
	"fmt"
	"time"
)

type Arg interface {
	int | string
}

func task[T Arg](id T) {
	fmt.Println(id)
	time.Sleep(time.Second * 2)
	fmt.Println("Task Completed", id)
}

func main() {
	//for i := 0; i < 10; i++ {
	//	go task(i)
	//}
	go task("Hello")
	fmt.Println("World")
	time.Sleep(time.Second * 2)
}
