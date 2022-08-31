package facade

import "fmt"

type PSU struct{}

func (PSU) turnOn() {
	fmt.Println("Power supply giving power")
}

type Fan struct{}

func (Fan) spinUp() {
	fmt.Println("Spinning up fans")
}

type PC struct {
	psu PSU
	fan Fan
}

func (p PC) turnOn() {
	p.psu.turnOn()
	p.fan.spinUp()
	fmt.Println("Now initializing other components")
}
