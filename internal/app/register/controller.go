package register

import (
	"fmt"
	"net/http"
)

type Reg struct {
	Email    string
	Password string
}

func Controller(w http.ResponseWriter, r *http.Request) {
	// inputs
	reg := Reg{
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}

	// validation

	var valid bool

	valid = IsEmailValid(reg.Email)

	fmt.Println("email: ", valid)

	valid = IsPasswordValid(reg.Password, 4)

	fmt.Println("pass: ", valid)

}
