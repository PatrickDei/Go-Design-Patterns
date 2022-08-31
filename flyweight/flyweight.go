package flyweight

import "fmt"

func ShowcaseFlyweight() {
	fmt.Println("\nFlyweight pattern")

	skin, _ := skinFactoryInstance.getSkinByType(T)
	fmt.Println(skin.getColor())

	skin, _ = skinFactoryInstance.getSkinByType(CT)
	fmt.Println(skin.getColor())

	skin, _ = skinFactoryInstance.getSkinByType(T)
	fmt.Println(skin.getColor())
}
