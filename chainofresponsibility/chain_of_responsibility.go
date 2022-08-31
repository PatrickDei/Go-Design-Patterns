package chainofresponsibility

import "fmt"

func ShowcaseChainOfResponsibility() {
	fmt.Println("\nChain of Responsibility pattern")
	h1 := ValidUsernameHandler{}
	h2 := UsernameExistsHandler{}
	h1.setNext(&h2)

	h1.handle("")
	h1.handle("Deisinger")
	h1.handle("Patrick")
}
