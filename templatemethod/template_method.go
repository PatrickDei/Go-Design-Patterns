package templatemethod

import "fmt"

func ShowcaseTemplateMethod() {
	fmt.Println("\nTemplate Method pattern")

	safeAction := Skeleton{changingBehavior: func() {
		fmt.Println("Custom action that does all the work being performed")
	}}

	safeAction.Run()
}
