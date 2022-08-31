package abstractfactory

type CorsairWorkSetupFactory struct{}

func (CorsairWorkSetupFactory) makeMouse() Mouse {
	return &CorsairMouse{
		MouseImpl{
			dpi:   12000,
			wired: false,
		},
	}
}

func (CorsairWorkSetupFactory) makeKeyboard() Keyboard {
	return &CorsairKeyboard{
		KeyboardImpl: KeyboardImpl{
			layout: "qwertz",
			keyNum: 25,
		},
		corsairLogoType: "standard",
	}
}
