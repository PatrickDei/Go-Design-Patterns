package chainofresponsibility

import "fmt"

type AuthHandler interface {
	handle(username string)
	setNext(AuthHandler)
}

type ValidUsernameHandler struct {
	nextHandler AuthHandler
}

func (v *ValidUsernameHandler) handle(username string) {
	if username == "" {
		fmt.Println("Invalid username")
		return
	}
	if v.nextHandler != nil {
		v.nextHandler.handle(username)
	}
}

func (v *ValidUsernameHandler) setNext(a AuthHandler) {
	v.nextHandler = a
}

type UsernameExistsHandler struct {
	nextHandler AuthHandler
}

func (u *UsernameExistsHandler) handle(username string) {
	if username == "Patrick" {
		fmt.Println("Yup, this account with name Patrick exists")
		return
	}
	fmt.Println("No such username")
	if u.nextHandler != nil {
		u.nextHandler.handle(username)
	}
}

func (u *UsernameExistsHandler) setNext(a AuthHandler) {
	u.nextHandler = a
}
