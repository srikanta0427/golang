package main

import "fmt"

type role int

const (
	roleUser role = iota
	roleAdmin
	roleModerator
)

type User struct {
	Name string
	Role role
}

func main() {
	user := User{
		Name: "John Doe",
		Role: roleAdmin,
	}
	if user.Role == roleAdmin {
		fmt.Println("Admin access granted")
	}
}
