package main

import (
	"fmt"
	"time"

	"github.com/Nimsaja/PatternsWithGo/tools"
)

var lastResult float64

func main() {
	stopTime(saveResult(tools.Pi, 10), 10)(10)
	fmt.Println("last Result ", lastResult)
}

func stopTime(a func(int) float64, n int) func(int) float64 {
	return func(n int) float64 {
		s := time.Now()
		r := a(n)
		fmt.Printf("calculate Pi inside stopTime %v -> %v\n", n, r)
		fmt.Println((time.Now()).Sub(s))

		return r
	}
}

func saveResult(a func(int) float64, n int) func(int) float64 {
	return func(n int) float64 {
		r := a(n)
		lastResult = r

		return r
	}
}
