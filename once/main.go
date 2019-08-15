package main

import (
	"fmt"
	"sync"
	"time"
)

var doOnce sync.Once

func main() {
	pts := 7

	for i := 1; i < 20; i++ {
		fmt.Printf("Calculation with %v is %v\n", i, calc(i, pts))
	}
}

func calc(n int, pts int) int {
	if n%3 == 0 {
		return initialize(n, pts)
	}
	return n
}

func initialize(n int, pts int) int {
	initFct := func(p int) {
		fmt.Printf("Initialize (%v)", n)
		for i := 0; i < p; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print(".")
		}
		fmt.Println()
	}

	doOnce.Do(func() {
		initFct(pts)
	})
	fmt.Printf("done (%v)\n", n)

	return n / 3
}
