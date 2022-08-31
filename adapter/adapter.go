package adapter

import "fmt"

func ShowcaseAdapter() {
	fmt.Println("\nAdapter pattern")
	fs := FruitServiceImpl{}

	df := DisplayableFruitsAdapter{
		FruitService: fs,
	}

	fmt.Println(df.getFruits())
}
