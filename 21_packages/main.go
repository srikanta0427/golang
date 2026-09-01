package main

import (
	"github.com/fatih/color"
	"github.com/srikanta/pck/auth/21_packages/auth"

	"github.com/srikanta/pck/auth/21_packages/user"
)

func main() {
	u := user.User{
		Name:     "root",
		Email:    "root@localhost",
		Password: "root",
		Username: "root123",
	}

	//auth.Login("root", "hash")
	color.Yellow(auth.SignUp(u).Email)
	color.Cyan("login")
	color.Red("login")
}
