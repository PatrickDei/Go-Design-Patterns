package abstractfactory

type LogitechWorkSetupFactory struct{}

func (LogitechWorkSetupFactory) makeMouse() Mouse {
	return &LogitechMouse{
		MouseImpl{
			dpi:   16000,
			wired: true,
		},
	}
}

func (LogitechWorkSetupFactory) makeKeyboard() Keyboard {
	return &LogitechKeyboard{
		KeyboardImpl: KeyboardImpl{
			layout: "qwerty",
			keyNum: 50,
		},
		logitechLogoType: "special logitech stuff",
	}
}
