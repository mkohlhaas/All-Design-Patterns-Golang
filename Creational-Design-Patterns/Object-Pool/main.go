package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

////////////////// Interface /////////////////////

type iPoolObject interface {
	// for comparing two different pool objects
	getID() string
}

/////////////// Implementation ///////////////////

type connection struct {
	id string
}

func (c *connection) getID() string {
	return c.id
}

////////////////// The Pool //////////////////////

type pool struct {
	idle     []iPoolObject
	active   []iPoolObject
	capacity int
	mulock   *sync.Mutex
}

func InitPool(poolObjects []iPoolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("Cannot create a pool of zero length.")
	}
	active := make([]iPoolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mulock:   new(sync.Mutex),
	}
	return pool, nil
}

/////////////// Pool API /////////////////////////

func (p *pool) Loan() (iPoolObject, error) {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	if len(p.idle) == 0 {
		return nil, fmt.Errorf("No pool object free. Please request after sometime.")
	}
	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)
	fmt.Printf("Loan Pool Object with ID: %s.\n", obj.getID())
	return obj, nil
}

func (p *pool) Receive(target iPoolObject) error {
	p.mulock.Lock()
	defer p.mulock.Unlock()
	err := p.remove(target)
	if err != nil {
		return err
	}
	p.idle = append(p.idle, target)
	fmt.Printf("Return Pool Object with ID: %s.\n", target.getID())
	return nil
}

// a little helper
func (p *pool) remove(target iPoolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("Targe pool object doesn't belong to the pool.")
}

//////////////////// Main ////////////////////////

func main() {
	connections := make([]iPoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}

	pool, err := InitPool(connections)
	if err != nil {
		log.Fatalf("Init Pool Error: %s.", err)
	}

	conn1, err := pool.Loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s.", err)
	}

	conn2, err := pool.Loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s.", err)
	}

	// ... do something with conn1, conn2

	// return conn1, conn2 to the pool
	pool.Receive(conn1)
	pool.Receive(conn2)
}
