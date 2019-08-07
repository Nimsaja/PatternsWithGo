package main

import (
	"fmt"
	"time"

	"github.com/Nimsaja/PatternsWithGo/tools"
)

var lastResult float64

type funcType = func(int) float64
type decType = func(funcType, int) funcType

func main() {
	decorators := []decType{stopTime, saveResult}

	a := 10
	decorate(tools.Pi, a, decorators...)(a)
	a = 10000
	decorate(tools.Pi, a, decorators...)(a)
	a = 1000000
	decorate(tools.Pi, a, decorators...)(a)
}

func decorate(f funcType, n int, ds ...decType) funcType {
	fct := f
	for _, d := range ds {
		fct = d(fct, n)
	}
	return fct
}

func stopTime(a funcType, n int) funcType {
	return func(n int) float64 {
		s := time.Now()
		r := a(n)
		fmt.Printf("calculate Pi inside stopTime %v -> %v\n", n, r)
		fmt.Println((time.Now()).Sub(s))

		return r
	}
}

func saveResult(a funcType, n int) funcType {
	return func(n int) float64 {
		r := a(n)
		lastResult = r
		fmt.Printf("save lastResult %v\n", r)

		return r
	}
}
