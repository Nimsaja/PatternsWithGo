package main

import (
	"fmt"
	"time"

	"github.com/Nimsaja/PatternsWithGo/tools"
)

var lastResult float64

type funcType = func(int) float64
type decType = func(funcType) funcType

func main() {
	decorators := []decType{stopTime, saveResult}

	decorate(tools.Pi, decorators...)(10)
	decorate(tools.Pi, decorators...)(10000)
	decorate(tools.Pi, decorators...)(1000000)

	decorate(tools.Fac, decorators...)(3)
	decorate(tools.Fac, decorators...)(6)
	decorate(tools.Fac, decorators...)(12)
}

func decorate(f funcType, ds ...decType) funcType {
	fct := f
	for _, d := range ds {
		fct = d(fct)
	}
	return fct
}

func stopTime(a funcType) funcType {
	return func(n int) float64 {
		s := time.Now()
		r := a(n)
		fmt.Printf("calculate Function inside stopTime %v -> %v\n", n, r)
		fmt.Println((time.Now()).Sub(s))

		return r
	}
}

func saveResult(a funcType) funcType {
	return func(n int) float64 {
		r := a(n)
		lastResult = r
		fmt.Printf("save lastResult %v\n", r)

		return r
	}
}
