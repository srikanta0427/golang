package main

import "fmt"

type status string

const (
	statusPending status = "pending"
	statusActive  status = "active"
	statusBlocked status = "blocked"
)

type Player struct {
	Name   string
	Status status
}

func (p Player) getData() {
	fmt.Println(p)
}

func main() {
	p := Player{
		Name:   "John Doe",
		Status: statusPending,
	}
	p.getData()
}
