package adapter

import "strings"

type DisplayableFruitsAdapter struct {
	FruitService
}

func (d DisplayableFruitsAdapter) getFruits() string {
	return strings.Join(d.FruitService.getFruits(), ", ")
}
