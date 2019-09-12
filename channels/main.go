package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int, 4)

	for i := 0; i < 100; i++ {
		c <- i
		go func(j int) {
			fmt.Printf("start with %v\n", j)
			doSomething(j)
			fmt.Printf("done with %v\n", <-c)
		}(i)
	}
}

func doSomething(i int) {
	r := rand.Float32() * 1000

	fmt.Printf("%v: sleep for %v seconds\n", i, r/1000)

	time.Sleep(time.Millisecond * time.Duration(r))
}
