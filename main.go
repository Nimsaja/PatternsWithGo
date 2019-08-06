package main

import (
	"fmt"
	"time"

	"github.com/Nimsaja/PatternsWithGo/tools"
)

var lastResult float64

func main() {
	stopTime(tools.Pi, 10)
	stopTime(tools.Pi, 10000)
	stopTime(tools.Pi, 1000000)

	saveResult(tools.Pi, 10)
	fmt.Println("last Result ", lastResult)
	saveResult(tools.Pi, 10000)
	fmt.Println("last Result ", lastResult)
	saveResult(tools.Pi, 1000000)
	fmt.Println("last Result ", lastResult)
}

func stopTime(a func(int) float64, n int) {
	s := time.Now()
	r := a(n)
	fmt.Printf("calculate Pi inside stopTime %v -> %v\n", n, r)
	fmt.Println((time.Now()).Sub(s))
}

func saveResult(a func(int) float64, n int) {
	r := a(n)
	lastResult = r
}
