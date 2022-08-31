package command

import "fmt"

func ShowcaseCommand() {
	fmt.Println("\nCommand pattern")
	tv := TV{isRunning: false}

	onCommand := OnCommand{
		device: &tv,
	}

	offCommand := OffCommand{
		device: &tv,
	}

	onButton := Button{
		command: &onCommand,
	}

	offButton := Button{
		command: &offCommand,
	}

	onButton.press()
	offButton.press()
	onButton.press()
	offButton.press()
}
