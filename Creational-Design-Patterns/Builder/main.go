package main

import "fmt"

/////////////////// Our Product //////////////////

type house struct {
	windowType string
	doorType   string
	floor      int
}

///////////////////// Interface /////////////////

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

//////////////// Builders ////////////////////////

//////// Normal House Builder ////////////////////

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *normalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *normalBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

/////////////// Igloo Builder ////////////////////

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *iglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *iglooBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

///////////////// House Types ////////////////////

type House interface {
	isHouse()
}

type houseType struct {
	sz string
}

func (houseType) isHouse() {}

func (h houseType) String() string {
	return h.sz
}

var (
	NormalHouse House = houseType{"Normal House"}
	Igloo       House = houseType{"Igloo"}
)

/////////////////// Public API ///////////////////

func NewBuilder(houseType House) (builder iBuilder) {
	switch houseType {
	case NormalHouse:
		builder = newNormalBuilder()
	case Igloo:
		builder = newIglooBuilder()
	}
	return
}

type director struct {
	builder iBuilder
}

func NewDirector() *director {
	return &director{}
}

func (d *director) BuildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func (d *director) SetBuilder(b iBuilder) {
	d.builder = b
}

///////////////// Main //////////////////////////

func main() {
	// Director
	director := NewDirector()

	// Normal House
	director.SetBuilder(NewBuilder(NormalHouse))
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n\n", normalHouse.floor)

	// Igloo
	director.SetBuilder(NewBuilder(Igloo))
	iglooHouse := director.BuildHouse()

	fmt.Printf("Igloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}
