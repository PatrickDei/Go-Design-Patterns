package abstractfactory

type Keyboard interface {
	getLayout() string
}

type KeyboardImpl struct {
	layout string
	keyNum int
}

func (k *KeyboardImpl) getLayout() string {
	return k.layout
}

// implementations

type LogitechKeyboard struct {
	KeyboardImpl
	logitechLogoType string
}
type CorsairKeyboard struct {
	KeyboardImpl
	corsairLogoType string
}
