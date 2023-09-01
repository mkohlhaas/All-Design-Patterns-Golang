package main

import (
	"fmt"
	"math"
)

// Separation of data and algorithms:
// Both are interfaces:
// data = shape
// processing = visitor

// Golang can accept interfaces as functions params,
// but it can't be the receiver of a method.

// Does not work (uncoment to see for yourself):
// func (i *interface{}) foo() {}

// data accepts visitors.
// visitors visit data.

// SO -> OS
// subject becomes object and object becomes subject

// `accept` and `visit` are symmetrical mirror images.
// Functions (processing) are data and vice versa.
// But the symmetry is broken by the data calling the very specific functions:
// - visitForSquare
// - visitForCircle
// - visitForRectangle

//////////// Interface for Shapes ////////////////

type shape interface {
	accept(visitor)
}

//////////// Visitor Interface ///////////////////

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForRectangle(*rectangle)
}

//////////////// Square //////////////////////////

type square struct {
	side float64
}

// The shape knows which visit method to call!!!
// subject square becomes object (func param) and
// object visitor becomes subject.
func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) String() string {
	return "Square"
}

//////////////// Circle //////////////////////////

type circle struct {
	radius float64
}

// The shape knows which visit method to call!!!
// subject circle becomes object (func param) and
// object visitor becomes subject.
func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) String() string {
	return "Circle"
}

//////////////// Rectangle ///////////////////////

type rectangle struct {
	l float64
	b float64
}

// The shape knows which visit method to call!!!
// subject rectangle becomes object (func param) and
// object visitor becomes subject.
func (t *rectangle) accept(v visitor) {
	v.visitForRectangle(t)
}

func (t *rectangle) String() string {
	return "Rectangle"
}

////////////// Area Visitor //////////////////////

type areaCalculator struct {
	area float64
}

var _ visitor = &areaCalculator{}

func (a *areaCalculator) visitForSquare(s *square) {
	a.area = s.side * s.side
	fmt.Printf("Area for %s: %0.2f\n", s, a.area)
}

func (a *areaCalculator) visitForCircle(c *circle) {
	a.area = math.Pi * c.radius * c.radius
	fmt.Printf("Area for %s: %0.2f\n", c, a.area)
}

func (a *areaCalculator) visitForRectangle(r *rectangle) {
	a.area = r.b * r.l
	fmt.Printf("Area for %s: %0.2f\n", r, a.area)
}

// subject visitor becomes object (func param) and
// object shape becomes subject.
func (v *areaCalculator) visit(s shape) {
	s.accept(v)
}

////// Middle Coordinates Visitor ////////////////

type middleCoordinates struct {
	x float64
	y float64
}

var _ visitor = &middleCoordinates{}

func (mc *middleCoordinates) visitForSquare(s *square) {
	mc.x = s.side / 2.0
	mc.y = s.side / 2.0
	fmt.Printf("Middle point coordinates for %s: x = %0.2f, y = %0.2f\n", s, mc.x, mc.y)
}

func (mc *middleCoordinates) visitForCircle(c *circle) {
	mc.x = 0
	mc.y = 0
	fmt.Printf("Middle point coordinates for %s: x = %0.2f, y = %0.2f\n", c, mc.x, mc.y)
}

func (mc *middleCoordinates) visitForRectangle(r *rectangle) {
	mc.x = r.l / 2.0
	mc.y = r.b / 2.0
	fmt.Printf("Middle point coordinates for %s: x = %0.2f, y = %0.2f\n", r, mc.x, mc.y)
}

// object shape becomes subject.
func (v *middleCoordinates) visit(s shape) {
	s.accept(v)
}

////////////// Main //////////////////////////////

func main() {
	// create shapes
	square := &square{side: 2}
	circle := &circle{radius: 3}
	rectangle := &rectangle{l: 2, b: 3}

	// create calculators
	areaCalculator := &areaCalculator{}
	mcCalculator := &middleCoordinates{}

	fmt.Println("Calculate areas:")
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)
	fmt.Println("\nThe same but idiomatic.")
	areaCalculator.visit(square)
	areaCalculator.visit(circle)
	areaCalculator.visit(rectangle)

	fmt.Println("\nCalculate middle points:")
	square.accept(mcCalculator)
	circle.accept(mcCalculator)
	rectangle.accept(mcCalculator)
	fmt.Println("\nThe same but idiomatic.")
	mcCalculator.visit(square)
	mcCalculator.visit(circle)
	mcCalculator.visit(rectangle)

	// areaCalculator
	fmt.Println("\nLooping over shapes:")
	shapes := []shape{square, circle, rectangle}
	for _, shape := range shapes {
		shape.accept(areaCalculator)
	}
	fmt.Println("\nThe same but idiomatic.")
	for _, shape := range shapes {
		areaCalculator.visit(shape)
	}

	// middleCoordinates
	fmt.Println("\nLooping over shapes:")
	for _, shape := range shapes {
		shape.accept(mcCalculator)
	}
	fmt.Println("\nThe same but idiomatic.")
	for _, shape := range shapes {
		mcCalculator.visit(shape)
	}
}
