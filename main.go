package main

import (
	"fmt"
	"time"

	"github.com/Nimsaja/PatternsWithGo/tools"
)

func main() {
	stopTime(tools.Pi, 10)
	stopTime(tools.Pi, 10000)
	stopTime(tools.Pi, 1000000)
}

func stopTime(a func(int) float64, n int) {
	s := time.Now()
	r := a(n)
	fmt.Printf("calculate Pi inside stopTime %v -> %v\n", n, r)
	fmt.Println((time.Now()).Sub(s))
}
