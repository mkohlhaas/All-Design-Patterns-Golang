package main

import "fmt"

//////////////// Interfaces //////////////////////

// Observer
type observer interface {
	update(string)
	getID() string
}

// Observee
type subject interface {
	register(Observer observer)
	deregister(Observer observer)
	notifyAll()
}

////////////// Interface implementations /////////

/////////////////// Observee /////////////////////

type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock.\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromslice(observerList []observer, observerToRemove observer) []observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

/////////////////// Observer /////////////////////

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s.\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}

//////////////////// Main ////////////////////////

func main() {
  // create observee
	shirtItem := newItem("Nike Shirt")
  // create observers
	observerFirst := &customer{id: "abc@gmail.com"}
	observerSecond := &customer{id: "xyz@gmail.com"}
  // register observers with observee
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

  // notify observers that shirts are available again
	shirtItem.updateAvailability()
}
