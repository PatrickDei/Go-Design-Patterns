package strategy

import "fmt"

func ShowcaseStrategy() {
	fmt.Println("\nStrategy pattern")

	externalProp := "basic"

	user := User{
		username: "john",
		password: "johhan",
	}

	var strategy TokenGenerationStrategy
	if externalProp == "basic" {
		strategy = &BasicTokenGenerationStrategy{}
	}
	if externalProp == "jwt" {
		strategy = &JWTTokenGenerationStrategy{}
	}

	fmt.Printf("Received token: %v\n", strategy.generateToken(user))
}
