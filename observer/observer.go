package observer

import "fmt"

func ShowcaseObserver() {
	fmt.Println("\nObserver pattern")
	userObservable := UserSubject{
		name:     "John",
		surname:  "Doe",
		username: "johndoe",
	}

	userObserver := UserObserverImpl{}

	userObservable.subscribe(&userObserver)

	userObservable.setName("Johnny")
}
