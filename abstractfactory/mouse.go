package abstractfactory

type Mouse interface {
	getDPI() int
}

type MouseImpl struct {
	dpi   int
	wired bool
}

func (m *MouseImpl) getDPI() int {
	return m.dpi
}

// implementations

type LogitechMouse struct {
	MouseImpl
}
type CorsairMouse struct {
	MouseImpl
}
