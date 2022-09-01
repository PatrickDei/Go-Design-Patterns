package templatemethod

import "fmt"

type Skeleton struct {
	changingBehavior func()
}

func (s *Skeleton) Run() {
	fmt.Println("Performing actions before the specified behavior")
	s.changingBehavior()
	fmt.Println("Cleanup after the action")
}
