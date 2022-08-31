package flyweight

// we save memory by only instantiating one of each skin instead of instantiating them for every player

type Skin interface {
	getColor() string
}

type TSkin struct{}

func (TSkin) getColor() string {
	return "brown"
}

type CTSkin struct{}

func (CTSkin) getColor() string {
	return "blue"
}
