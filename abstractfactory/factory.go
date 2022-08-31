package abstractfactory

type WorkSetupFactory interface {
	makeMouse() Mouse
	makeKeyboard() Keyboard
}
