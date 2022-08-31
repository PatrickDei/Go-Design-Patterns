package factorymethod

type Color interface {
	getHexColor() string
}

type Red struct{}

func (Red) getHexColor() string {
	return "#ff0000"
}

type Green struct{}

func (Green) getHexColor() string {
	return "#00ff00"
}

type Blue struct{}

func (Blue) getHexColor() string {
	return "#0000ff"
}
