package main

import "fmt"

type Paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway Paymenter
	//gateway razorpay
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("razorpay pay", amount)
}

func main() {
	rz := razorpay{}
	p := payment{
		gateway: rz,
	}
	p.makePayment(100)
}
