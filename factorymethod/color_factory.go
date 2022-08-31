package factorymethod

import (
	"fmt"
	"os"
)

type ColorFactory struct{}

func (ColorFactory) createColor(colorName string) Color {
	switch colorName {
	case "Red":
		return Red{}
	case "Blue":
		return Blue{}
	case "Green":
		return Green{}
	default:
		fmt.Println("Not a recognized color")
		os.Exit(-1)
		return nil
	}
}
