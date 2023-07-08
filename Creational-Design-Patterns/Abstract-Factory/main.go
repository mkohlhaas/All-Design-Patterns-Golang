package main

import "fmt"

///////////////////// Interfaces ////////////////

type iShoe interface {
	getLogo() string
	getSize() int
}

type iShort interface {
	getColor() string
	getMeasure() int
}

type iSportsFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

//////////////////// Factories ///////////////////

//////////////// Adidas Factory //////////////////

type adidasFactory struct{}

func (a *adidasFactory) makeShoe() iShoe {
	return &adidasShoe{shoe{logo: "adidas", size: 14}}
}

func (a *adidasFactory) makeShort() iShort {
	return &adidasShort{short{color: "red", measure: 15}}
}

//////////////////// Nike Factory ////////////////

type nikeFactory struct{}

func (n *nikeFactory) makeShoe() iShoe {
	return &nikeShoe{shoe{logo: "nike", size: 16}}
}

func (n *nikeFactory) makeShort() iShort {
	return &nikeShort{short{color: "blue", measure: 17}}
}

/////// Implementation of Shoe Interface /////////

type shoe struct {
	logo string
	size int
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) getSize() int {
	return s.size
}

/////// Implementation of Short Interface ////////

type short struct {
	color   string
	measure int
}

func (s *short) getColor() string {
	return s.color
}

func (s *short) getMeasure() int {
	return s.measure
}

//////////////////// Embedding ///////////////////

type adidasShoe struct {
	shoe
}

type nikeShoe struct {
	shoe
}

type adidasShort struct {
	short
}

type nikeShort struct {
	short
}

////////////////// Public API ////////////////////

type Brand interface {
	isBrand() // unexported method: cannot create new Brands
}

// quasi enumeration types (type-safe)
var (
	Adidas Brand = brand{"Adidas"}
	Nike   Brand = brand{"Nike"}
)

func NewSportsFactory(brand Brand) (sf iSportsFactory) {
	switch brand {
	case Adidas:
		sf = &adidasFactory{}
	case Nike:
		sf = &nikeFactory{}
	}
	return
}

///////////////////// Helpers ////////////////////

type brand struct{ sz string } // unexported type: cannot create new brands

func (brand) isBrand() {} // implement the exported interface

func (b brand) String() string { return b.sz } // Stringer interface

////////////////////// Main //////////////////////

func main() {
	brands := []Brand{Adidas, Nike}

	for _, brand := range brands {
		factory := NewSportsFactory(brand)
		fmt.Println("--- ", brand, " ---")
		printShoeDetails(factory.makeShoe())
		printShortDetails(factory.makeShort())
	}
}

func printShoeDetails(s iShoe) {
	fmt.Println("Shoe")
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n\n", s.getSize())
}

func printShortDetails(s iShort) {
	fmt.Println("Short")
	fmt.Printf("Color: %s\n", s.getColor())
	fmt.Printf("Measure: %d\n\n", s.getMeasure())
}
