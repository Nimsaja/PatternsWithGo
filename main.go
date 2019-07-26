package main

import (
	"fmt";
	"math"
)

func main() {
	calcPi(100000)
	calcPi(10000000)
}

func calcPi(n int) {
	//Sum[1:oo](1/n^2)=pi^2/6
	s := 0.0
	for i:=1; i<=n; i++ {
		s = s + 1/math.Pow(float64(i), 2)
	}
	pi := math.Sqrt(6.0*s)
	
	fmt.Println("Pi is ", pi)
}