package abstractfactory

import "fmt"

func ShowcaseAbstractFactory() {
	fmt.Println("\nAbstract Factory pattern")
	corsFactory := CorsairWorkSetupFactory{}
	logiFactory := LogitechWorkSetupFactory{}

	logiMouse := logiFactory.makeMouse()
	logiKeyboard := logiFactory.makeKeyboard()

	corsMouse := corsFactory.makeMouse()
	corsKeyboard := corsFactory.makeKeyboard()

	fmt.Printf("%+v\n", logiMouse)
	fmt.Printf("%+v\n", corsMouse)
	fmt.Printf("%+v\n", logiKeyboard)
	fmt.Printf("%+v\n", corsKeyboard)
}
