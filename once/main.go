package main

import (
	"fmt"
	"sync"
	"time"
)

var doOnce sync.Once

func main() {
	for i := 1; i < 20; i++ {
		fmt.Printf("Calculation with %v is %v\n", i, calc(i))
	}
}

func calc(n int) int {
	if n%3 == 0 {
		return initialize(n)
	}
	return n
}

func initialize(n int) int {
	doOnce.Do(func() {
		fmt.Printf("Initialize (%v)", n)
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Print(".")
		}
		fmt.Println()
	})
	fmt.Printf("done (%v)\n", n)

	return n / 3
}
