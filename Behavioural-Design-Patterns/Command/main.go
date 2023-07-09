package main

import "fmt"

///////////////// Interfaces /////////////////////

type command interface {
	execute()
}

type device interface {
	on()
	off()
}

///////////////// Commands ///////////////////////

/////////////// On Command ///////////////////////

// command embeds receiver
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

/////////////// Off Command //////////////////////

// command embeds receiver
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

////////////// Invoker ///////////////////////////

// invoker embeds command
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

///////////////// Receiver ///////////////////////

// tv can be turned on and off
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on.")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off.")
}

//////////////// Main ////////////////////////////

func main() {
  // create receiver
	tv := &tv{}

  // create on command
	onCommand := &onCommand{
		device: tv,
	}

  // create off command
	offCommand := &offCommand{
		device: tv,
	}

  // create invoker for turning on
	onButton := &button{
		command: onCommand,
	}

  // create invoker for turning off
	offButton := &button{
		command: offCommand,
	}

  // invoke commands
	onButton.press()
	offButton.press()
}
