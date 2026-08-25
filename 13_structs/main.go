package main

import (
	"fmt"
	"time"
)

// creating structs
type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

// receiver type
func (o *Order) changedStatus(status string) {
	o.status = status
}

func changed(order *Order, status string) {
	order.status = status
}

// constructor
func newOrder(id string, amount float32, status string) *Order {
	order := Order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &order
}

func main() {
	order := Order{
		id:     "123",
		amount: 1.2,
		status: "pending",
		//createdAt: time.Now(),
	}
	order.createdAt = time.Now()
	order.changedStatus("confirmed")
	fmt.Println(order)
	changed(&order, "pending")
	fmt.Println(order)

	// use constructor
	myOrder := newOrder("123", 1.2, "pending")
	myOrder.createdAt = time.Now()
	fmt.Println(myOrder)

	// creating structs in another way
	Book := struct {
		id     string
		amount float32
	}{id: "123", amount: 1.2}
	fmt.Println(Book)
}
