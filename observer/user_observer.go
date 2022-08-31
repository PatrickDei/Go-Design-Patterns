package observer

import "fmt"

type UserObserver interface {
	onChange(UserSubject)
}

type UserObserverImpl struct{}

func (uo *UserObserverImpl) onChange(us UserSubject) {
	fmt.Printf("%+v\n", us)
}
