package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	price   float64
}

type EBook struct {
	e_type string
	Book
}

func main() {
	myEBook1 := EBook{
		e_type:  "pdf",
		author:  "wang xiao",
		price:   1.2,
		title:   "s",
		subject: "m",
	}

	myBook := Book{
		title:   "Go Programming Language",
		author:  "Go Dev",
		subject: "Programming Language",
		price:   1.2,
	}

	myEBook3 := EBook{
		e_type: "pdf",
		Book: Book{
			title:   "Go Programming Language",
			author:  "Go Dev",
			subject: "Programming Language",
			price:   1.2,
		},
	}

	myEBook2 := EBook{
		e_type: "pdf",
		Book:   myBook,
	}
	fmt.Println(myEBook1, myEBook2, myEBook3)
}
