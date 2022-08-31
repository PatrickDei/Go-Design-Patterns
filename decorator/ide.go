package decorator

import "fmt"

type IDE interface {
	turnOn()
}

type Intellij struct{}

func (i *Intellij) turnOn() {
	fmt.Println("Turning on. Taking you back to where we left off.")
}

type FriendlyIntellij struct {
	intellij Intellij
}

func (i *FriendlyIntellij) turnOn() {
	fmt.Println("Hope you have a great day!")
	i.intellij.turnOn()
}

type PassiveAggressiveIntellij struct {
	intellij Intellij
}

func (i *PassiveAggressiveIntellij) turnOn() {
	fmt.Println("Again? -.-")
	i.intellij.turnOn()
}
