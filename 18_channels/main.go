package main

import (
	"fmt"
	"time"
)

func processNum(msg chan<- string) {
	result := "some result"
	time.Sleep(time.Second * 2)
	msg <- result
}

func main() {

	msgChain := make(chan string)
	go processNum(msgChain)
	msg := <-msgChain
	fmt.Println(msg)
	fmt.Println("work after receiving msg(chain)")

}
