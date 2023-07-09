package main

import "fmt"

////////// Abstraction Interface /////////////////

type computer interface {
	print()
	setPrinter(printer)
}

/////////////// Computers ////////////////////////

///////////////// Mac ////////////////////////////

// abstraction uses implementation via a bridge (embedding)
type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print request for mac.")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

/////////////// Windows //////////////////////////

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("Print request for windows.")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

///////// Implementation Interface ///////////////

type printer interface {
	printFile()
}

////////////////// Printers //////////////////////

/////////////////// Epson ////////////////////////

type epson struct{}

func (p *epson) printFile() {
	fmt.Println("Printing by an EPSON Printer.")
}

//////////////////// HP //////////////////////////

type hp struct{}

func (p *hp) printFile() {
	fmt.Println("Printing by an HP Printer.")
}

///////////////// Main ///////////////////////////

func main() {
	// create computers
	macComputer := &mac{}
	winComputer := &windows{}

	// create printers
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	// print on Mac with HP
	macComputer.setPrinter(hpPrinter)
	macComputer.print()

	// print on Mac with Epson
	macComputer.setPrinter(epsonPrinter)
	macComputer.print()

	// print on Windows with HP
	winComputer.setPrinter(hpPrinter)
	winComputer.print()

	// print on Windows with Epson
	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
}
