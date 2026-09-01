package auth

import (
	"fmt"

	"github.com/srikanta/pck/auth/21_packages/user"
)

func lgoinWithCredential(username, password string) {
	fmt.Print("LoginWithCredential", username, password)
}

func Login(username, password string) {
	fmt.Print("Login", username, password)
}

func SignUp(user user.User) user.User {
	return user
}
