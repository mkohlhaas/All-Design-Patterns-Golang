package main

import "fmt"

///////////// Interface //////////////////////////

type computer interface {
	insertInSquarePort()
}

////// Interface Implementations /////////////////

///////////////// Mac ////////////////////////////

type mac struct {
}

func (m *mac) insertInSquarePort() {
	fmt.Println("Inserting square port into mac machine.")
}

/////////// Windows Adapter //////////////////////

// embeds Windows!
type windowsAdapter struct {
	windowMachine *windows
}

// adapter just forwards the call
func (w *windowsAdapter) insertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}

////////////// Windows ///////////////////////////

// NOT a computer interface implementation!
type windows struct{}

func (w *windows) insertInCirclePort() {
	fmt.Println("Inserting circle port into windows machine.")
}

//////////////// Client //////////////////////////

// client doesn't care about different port types
type client struct {
}

func (c *client) insertSquareUsbInComputer(com computer) {
	com.insertInSquarePort()
}

/////////////////// Main /////////////////////////

func main() {
	// create a client
	client := &client{}

	// connect to Mac
	mac := &mac{}
	client.insertSquareUsbInComputer(mac)

	// connect to windows via adapter
	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowsMachine, // embedding the windows machine
	}
	client.insertSquareUsbInComputer(windowsMachineAdapter)
}
