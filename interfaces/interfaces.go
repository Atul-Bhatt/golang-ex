package main

import (
	"fmt"
	"math"
)

// Interface geometry
type geometry interface {
	area() float64
	per() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width float64
	height float64
}

// Implement geometry for circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) per() float64 {
	return 2 * math.Pi * c.radius
}

// Implement geometry for rect
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) per() float64 {
	return 2*r.width + 2*r.height
}

// Function that accepts value of type that implements geometry
func measure(shape geometry) {
	fmt.Println(shape.area())
	fmt.Println(shape.per())
}

func main() {
	r := rect{width: 2, height: 10}
	c := circle{radius: 10}

	measure(r)
	measure(c)
}
