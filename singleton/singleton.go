package singleton

import "fmt"

func ShowcaseSingleton() {
	fmt.Println("\nSingleton pattern")

	for i := 0; i < 10; i++ {
		go getInstance()
	}

	for i := 0; i < 10; i++ {
		go getAlternativeInstance()
	}

	// pause execution
	// fmt.Scanln()
}
