package adapter

type FruitService interface {
	getFruits() []string
}

type FruitServiceImpl struct{}

func (FruitServiceImpl) getFruits() []string {
	return []string{"Apple", "Orange", "Pear"}
}
