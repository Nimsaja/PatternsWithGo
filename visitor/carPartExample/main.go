package main

import "fmt"

// this is an example from the web

// CarPart defines the common interface for all elements that should be visitable
type CarPart interface {
	Accept(CarPartVisitor)
}

// Wheel and Engine are Concrete CarParts. The Accept methods call the according visitXXX method of the visitor
// and pass a reference of themselves
type Wheel struct {
	Name string
}

func (this *Wheel) Accept(visitor CarPartVisitor) {
	visitor.visitWheel(this)
}

// Engine struct and function
type Engine struct{}

func (this *Engine) Accept(visitor CarPartVisitor) {
	visitor.visitEngine(this)
}

// Car is the Object Structure maintaining a list of CarPart objects
type Car struct {
	parts []CarPart
}

// NewCar assembles a car this with four wheels and an engine
func NewCar() *Car {
	this := new(Car)
	this.parts = []CarPart{
		&Wheel{"front left"},
		&Wheel{"front right"},
		&Wheel{"rear left"},
		&Wheel{"rear right"},
		&Engine{}}
	return this
}

// Accept method which iterates over the carParts and calls their Accept method
func (this *Car) Accept(visitor CarPartVisitor) {
	for _, part := range this.parts {
		part.Accept(visitor)
	}
}

// CarPartVisitor is the common Visitor interface, declaring visitXXX operations, one for each concrete element type
type CarPartVisitor interface {
	visitWheel(wheel *Wheel)
	visitEngine(engine *Engine)
}

// PrintVisitor is a Concrete Visitor
type PrintVisitor struct{}

func (this *PrintVisitor) visitWheel(wheel *Wheel) {
	fmt.Printf("Visiting the %v wheel\n", wheel.Name)
}

func (this *PrintVisitor) visitEngine(engine *Engine) {
	fmt.Printf("Visiting engine\n")
}

// *******main************
func main() {
	car := NewCar()
	car.Accept(new(PrintVisitor))
}
