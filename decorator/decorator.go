package decorator

import "fmt"

func ShowcaseDecorator() {
	fmt.Println("\nDecorator pattern")
	intellij := Intellij{}

	friendlyIntellij := FriendlyIntellij{intellij: intellij}

	paIntellij := PassiveAggressiveIntellij{intellij: intellij}

	intellij.turnOn()
	friendlyIntellij.turnOn()
	paIntellij.turnOn()
}
