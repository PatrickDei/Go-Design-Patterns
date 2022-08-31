package bridge

import "fmt"

type Computer interface {
	Print()
	SetPrinter(Printer)
}

// implementations

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Machine: Mac")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Machine: Windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
