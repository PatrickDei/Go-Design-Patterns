package facade

import "fmt"

func ShowcaseFacade() {
	fmt.Println("\nFacade pattern")
	pc := PC{
		psu: PSU{},
		fan: Fan{},
	}

	pc.turnOn()
}
