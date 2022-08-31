package bridge

import "fmt"

func ShowcaseBridge() {
	fmt.Println("\nBridge pattern")
	hpPrinter := HP{}
	samsungPrinter := Samsung{}

	mac := Mac{printer: hpPrinter}
	mac.Print()

	mac.SetPrinter(samsungPrinter)
	mac.Print()

	windows := Windows{printer: hpPrinter}
	windows.Print()

	windows.SetPrinter(samsungPrinter)
	windows.Print()
}
