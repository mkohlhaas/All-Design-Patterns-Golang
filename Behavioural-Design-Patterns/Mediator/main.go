package main

import (
	"fmt"
	"sync"
)

////////////// Interfaces ////////////////////////

type train interface {
	requestArrival()
	departure()
	permitArrival()
}

type mediator interface {
	canLand(train) bool
	notifyFree()
}

//////////////// Trains //////////////////////////

//////////// Passenger Train /////////////////////

type passengerTrain struct {
	mediator mediator
}

func (g *passengerTrain) requestArrival() {
	if g.mediator.canLand(g) {
		fmt.Println("Passenger Train: Landing.")
	} else {
		fmt.Println("Passenger Train: Waiting.")
	}
}

func (g *passengerTrain) departure() {
	fmt.Println("Passenger Train: Leaving.")
	g.mediator.notifyFree()
}

func (g *passengerTrain) permitArrival() {
	fmt.Println("Passenger Train: Arrival Permitted. Landing.")
}

//////////// Goods Train /////////////////////////

type goodsTrain struct {
	mediator mediator
}

func (g *goodsTrain) requestArrival() {
	if g.mediator.canLand(g) {
		fmt.Println("Goods Train: Landing.")
	} else {
		fmt.Println("Goods Train: Waiting.")
	}
}

func (g *goodsTrain) departure() {
	g.mediator.notifyFree()
	fmt.Println("Goods Train: Leaving.")
}

func (g *goodsTrain) permitArrival() {
	fmt.Println("Goods Train: Arrival Permitted. Landing.")
}

////////////////// Mediator Implementation ///////

type stationManager struct {
	isPlatformFree bool
	lock           *sync.Mutex
	trainQueue     []train
}

func newStationManger() *stationManager {
	return &stationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *stationManager) canLand(t train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *stationManager) notifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.isPlatformFree = true
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

/////////////// Main /////////////////////////////

func main() {
  // create mediator
	stationManager := newStationManger()

  // create trains
	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}

	goodsTrain := &goodsTrain{
		mediator: stationManager,
	}

  // ... some action ...
	passengerTrain.requestArrival()
	goodsTrain.requestArrival()
	passengerTrain.departure()
}
