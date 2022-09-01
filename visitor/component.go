package visitor

type Component interface {
	getPrice() int
	getPower() int
}

type PSU struct{}

func (PSU) getPrice() int {
	return 100
}

func (PSU) getPower() int {
	return 750
}

type CPU struct{}

func (CPU) getPrice() int {
	return 1000
}

func (CPU) getPower() int {
	return 150
}
