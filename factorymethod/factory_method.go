package factorymethod

import "fmt"

func ShowcaseFactoryMethod() {
	fmt.Println("\nFactory Method pattern")

	cf := ColorFactory{}
	fmt.Println(cf.createColor("Red").getHexColor())
	fmt.Println(cf.createColor("Blue").getHexColor())
	fmt.Println(cf.createColor("Green").getHexColor())
}
