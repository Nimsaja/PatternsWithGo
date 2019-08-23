package main

import (
	"github.com/Nimsaja/PatternsWithGo/visitor/geometricForms/forms"
)

func main() {
	geom := make([]forms.GeomForm, 0)

	geom = append(geom, &forms.Square{A: 7.0})
	geom = append(geom, &forms.Rectangle{A: 7, B: 2})
	geom = append(geom, &forms.Trapeze{A: 7, B: 9, C: 5, D: 8, H: 7})
	geom = append(geom, &forms.Triangle{A: 7, B: 9, C: 5, H: 7})
	geom = append(geom, &forms.Circle{R: 7})

	pv := new(forms.PrintVisitor)
	a := new(forms.AreaVisitor)
	p := new(forms.PerimeterVisitor)
	for _, g := range geom {
		g.Accept(pv)
		g.Accept(a)
		g.Accept(p)
	}
}
