package forms

// GeomForm interface
type GeomForm interface {
	Accept(GeomFormVisitor)
}

// Square with side length a
type Square struct {
	A float32
}

// Accept ...
func (q *Square) Accept(visitor GeomFormVisitor) {
	visitor.visitSquare(q)
}

// Rectangle with side length a and b
type Rectangle struct {
	A float32
	B float32
}

// Accept ...
func (r *Rectangle) Accept(visitor GeomFormVisitor) {
	visitor.visitRectangle(r)
}

// Trapeze with side length a, b, c, d and height h
type Trapeze struct {
	A float32
	B float32
	C float32
	D float32
	H float32
}

// Accept ...
func (t *Trapeze) Accept(visitor GeomFormVisitor) {
	visitor.visitTrapeze(t)
}

// Triangle with side length a, b, c and height h
type Triangle struct {
	A float32
	B float32
	C float32
	H float32
}

// Accept ...
func (t *Triangle) Accept(visitor GeomFormVisitor) {
	visitor.visitTriangle(t)
}

// Circle with radius r
type Circle struct {
	R float32
}

// Accept ...
func (c *Circle) Accept(visitor GeomFormVisitor) {
	visitor.visitCircle(c)
}
