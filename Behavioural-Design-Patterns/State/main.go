package main

import (
	"fmt"
	"log"
)

/////////////////// Custom Error Type ////////////

type stateError struct {
	err error
}

func (e *stateError) Unwrap() {
	if e.err != nil {
		log.Fatalf(e.err.Error())
	}
}

/////////////////// Interface ////////////////////

// flow of state: request item → insert money → dispense item
// adding items shall only be allow in noItemState and hasItemState
type state interface {
	addItem(int) *stateError
	requestItem() *stateError
	insertMoney(money int) *stateError
	dispenseItem() *stateError
}

/////// States (Interface Implementations) ///////

////////////// No Item ///////////////////////////

type noItemState struct {
	vendingMachine *vendingMachine
}

func (i *noItemState) addItem(count int) *stateError {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	fmt.Printf("%d items added.\n", count)
	return &stateError{nil}
}

func (i *noItemState) requestItem() *stateError {
	return &stateError{fmt.Errorf("Item out of stock.")}
}

func (i *noItemState) insertMoney(money int) *stateError {
	return &stateError{fmt.Errorf("Item out of stock.")}
}
func (i *noItemState) dispenseItem() *stateError {
	return &stateError{fmt.Errorf("Item out of stock.")}
}

/////////////// Has Item /////////////////////////

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) addItem(count int) *stateError {
	i.vendingMachine.incrementItemCount(count)
	fmt.Printf("%d items added.\n", count)
	return &stateError{nil}
}

func (i *hasItemState) requestItem() *stateError {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return &stateError{fmt.Errorf("No item present.")}
	}
	fmt.Printf("Item requestd.\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return &stateError{nil}
}

func (i *hasItemState) insertMoney(money int) *stateError {
	return &stateError{fmt.Errorf("Please select item first.")}
}
func (i *hasItemState) dispenseItem() *stateError {
	return &stateError{fmt.Errorf("Please select item first.")}
}

////////////////// Item Requested ////////////////

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) addItem(count int) *stateError {
	return &stateError{fmt.Errorf("Item Dispense in progress.")}
}

func (i *itemRequestedState) requestItem() *stateError {
	return &stateError{fmt.Errorf("Item already requested.")}
}

func (i *itemRequestedState) insertMoney(money int) *stateError {
	if money < i.vendingMachine.itemPrice {
		return &stateError{fmt.Errorf("Inserted money is less. Please insert %d.", i.vendingMachine.itemPrice)}
	}
	fmt.Println("Money entered is ok.")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return &stateError{nil}
}

func (i *itemRequestedState) dispenseItem() *stateError {
	return &stateError{fmt.Errorf("Please insert money first.")}
}

/////////////// Has Money ////////////////////////

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (i *hasMoneyState) requestItem() *stateError {
	return &stateError{fmt.Errorf("Item dispense in progress.")}
}

func (i *hasMoneyState) addItem(count int) *stateError {
	return &stateError{fmt.Errorf("Item dispense in progress.")}
}

func (i *hasMoneyState) insertMoney(money int) *stateError {
	return &stateError{fmt.Errorf("Stop inserting money! Balance is already enough.")}
}

func (i *hasMoneyState) dispenseItem() *stateError {
	fmt.Println("Dispensing Item.")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return &stateError{nil}
}

/////////////////// Context //////////////////////

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state // is one of the four states above

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	// fill in the four predefined states
	v.hasItem = &hasItemState{vendingMachine: v}
	v.itemRequested = &itemRequestedState{vendingMachine: v}
	v.hasMoney = &hasMoneyState{vendingMachine: v}
	v.noItem = &noItemState{vendingMachine: v}

	// set initial state
	v.setState(v.hasItem)
	return v
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items.\n", count)
	v.itemCount = v.itemCount + count
}

// Context is also a state. Forwards methods to currentState.
func (v *vendingMachine) addItem(count int) *stateError {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) requestItem() *stateError {
	return v.currentState.requestItem()
}

func (v *vendingMachine) insertMoney(money int) *stateError {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() *stateError {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

////////////////// Main //////////////////////////

func main() {
	vendingMachine := newVendingMachine(1, 10)
	vendingMachine.requestItem().Unwrap()
	vendingMachine.insertMoney(10).Unwrap()
	vendingMachine.dispenseItem().Unwrap()
	vendingMachine.addItem(2).Unwrap()
	vendingMachine.requestItem().Unwrap()
	vendingMachine.insertMoney(10).Unwrap()
	vendingMachine.dispenseItem().Unwrap()
}
