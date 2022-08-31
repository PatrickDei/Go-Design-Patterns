package bridge

import "fmt"

type Printer interface {
	PrintFile()
}

// implementations

type HP struct{}

func (HP) PrintFile() {
	fmt.Println("Printer: HP")
}

type Samsung struct{}

func (Samsung) PrintFile() {
	fmt.Println("Printer: Samsung")
}
