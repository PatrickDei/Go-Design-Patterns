package command

import "fmt"

type Device interface {
	on()
	off()
}

type TV struct {
	isRunning bool
}

func (t *TV) on() {
	t.isRunning = true
	fmt.Println("TV turned on")
}

func (t *TV) off() {
	t.isRunning = false
	fmt.Println("TV turned off")
}

type OnCommand struct {
	device Device
}

func (o *OnCommand) execute() {
	o.device.on()
}

type OffCommand struct {
	device Device
}

func (o *OffCommand) execute() {
	o.device.off()
}

type Command interface {
	execute()
}

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}
