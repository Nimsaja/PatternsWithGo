package main

import (
	"fmt"
	"math"
	"time"
)

var cache map[int]float64

func init() {
	cache = make(map[int]float64)
}

type piFunc func(int) float64

func pi(n int) float64 {
	s := 0.0
	for i := 1; i <= n; i++ {
		s = s + 1/math.Pow(float64(i), 2)
	}
	pi := math.Sqrt(6.0 * s)

	return pi
}

func stop(f piFunc) piFunc {
	return func(n int) float64 {
		s := time.Now()
		r := f(n)
		fmt.Printf("calculate Pi inside stopTime %v -> %v\n", n, r)
		fmt.Printf("time: %s\n", time.Since(s))
		return r
	}
}

func fromCache(f piFunc) piFunc {
	return func(n int) float64 {
		r, ok := cache[n]
		if !ok {
			r = f(n)
			cache[n] = r
		}
		return r
	}
}

func main() {
	n := 100000

	stop(pi)(n)

	stop(fromCache(pi))(n)
	stop(fromCache(pi))(n)
}
