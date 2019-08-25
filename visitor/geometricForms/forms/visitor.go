package forms

import (
	"fmt"
	"math"
)

// GeomFormVisitor interface
type GeomFormVisitor interface {
	visitSquare(q *Square)
	visitRectangle(r *Rectangle)
	visitTrapeze(t *Trapeze)
	visitTriangle(t *Triangle)
	visitCircle(c *Circle)
}

// PrintVisitor is a Concrete Visitor to print out the properties of a geometric form
type PrintVisitor struct{}

func (prv *PrintVisitor) visitSquare(q *Square) {
	fmt.Printf("\nVisiting Square a = %v\n", q.A)
}

func (prv *PrintVisitor) visitRectangle(r *Rectangle) {
	fmt.Printf("\nVisiting Rectangle a = %v, b = %v\n", r.A, r.B)
}

func (prv *PrintVisitor) visitTrapeze(t *Trapeze) {
	fmt.Printf("\nVisiting Trapeze a = %v, b = %v, c = %v, d = %v, h = %v\n", t.A, t.B, t.C, t.D, t.H)
}

func (prv *PrintVisitor) visitTriangle(t *Triangle) {
	fmt.Printf("\nVisiting Triangle a = %v, b = %v, c = %v, h = %v\n", t.A, t.B, t.C, t.H)
}

func (prv *PrintVisitor) visitCircle(c *Circle) {
	fmt.Printf("\nVisiting Circle r = %v\n", c.R)
}

// AreaVisitor is a Concrete Visitor to calculate the Area of a geometric form
type AreaVisitor struct{}

func (av *AreaVisitor) visitSquare(q *Square) {
	fmt.Printf("Area = %v\n", q.A*q.A)
}

func (av *AreaVisitor) visitRectangle(r *Rectangle) {
	fmt.Printf("Area = %v\n", r.A*r.B)
}

func (av *AreaVisitor) visitTrapeze(t *Trapeze) {
	fmt.Printf("Area = %v\n", (t.A+t.C)*t.H/2)
}

func (av *AreaVisitor) visitTriangle(t *Triangle) {
	fmt.Printf("Area = %v\n", (t.C*t.H)/2)
}

func (av *AreaVisitor) visitCircle(c *Circle) {
	fmt.Printf("Area = %v\n", math.Pi*c.R*c.R)
}

// PerimeterVisitor is a Concrete Visitor to calculate the Perimeter of a geometric form
type PerimeterVisitor struct{}

func (pv *PerimeterVisitor) visitSquare(q *Square) {
	fmt.Printf("Perimeter = %v\n", 4*q.A)
}

func (pv *PerimeterVisitor) visitRectangle(r *Rectangle) {
	fmt.Printf("Perimeter = %v\n", 2*(r.A+r.B))
}

func (pv *PerimeterVisitor) visitTrapeze(t *Trapeze) {
	fmt.Printf("Perimeter = %v\n", t.A+t.B+t.C+t.D)
}

func (pv *PerimeterVisitor) visitTriangle(t *Triangle) {
	fmt.Printf("Perimeter = %v\n", t.A+t.B+t.C)
}

func (pv *PerimeterVisitor) visitCircle(c *Circle) {
	fmt.Printf("Perimeter = %v\n", 2*math.Pi*c.R)
}
