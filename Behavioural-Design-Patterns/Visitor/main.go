package main

import (
	"fmt"
	"math"
)

//////////// Interface for Shapes ////////////////

type shape interface {
	accept(visitor)
}

//////////////// Square //////////////////////////

type square struct {
	side float64
}

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

func (t *rectangle) accept(v visitor) {
	v.visitForRectangle(t)
}

func (t *rectangle) String() string {
	return "Rectangle"
}

//////////// Visitor Interface ///////////////////

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForRectangle(*rectangle)
}

////////////// Area Visitor //////////////////////

type areaCalculator struct {
	area float64
}

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

////// Middle Coordinates Visitor ////////////////

type middleCoordinates struct {
	x float64
	y float64
}

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

////////////// Main //////////////////////////////

func main() {
  // create shapes
	square := &square{side: 2}
	circle := &circle{radius: 3}
	rectangle := &rectangle{l: 2, b: 3}

  // create area calculator
	areaCalculator := &areaCalculator{}

  // calculate areas
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

  // create middle coordinates calculator
	mcCalculator := &middleCoordinates{}
	square.accept(mcCalculator)
	circle.accept(mcCalculator)
	rectangle.accept(mcCalculator)
}
