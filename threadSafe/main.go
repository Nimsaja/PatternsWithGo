package main

import (
	"fmt"
	"sync"
	"time"
)

type Number struct {
	sync.RWMutex
	n int
}

var nb = new(Number)

func main() {
	go safeIncrease()
	printOut("1")
	go safeIncrease()
	printOut("2")
	go safeIncrease()
	printOut("3")
	go safeIncrease()
	printOut("4")
	go safeIncrease()
	printOut("5")

	time.Sleep(time.Second * 1)

	printOut("FIN")
}

func safeIncrease() {
	nb.Lock()
	defer nb.Unlock()
	nb.increase()
}

func printOut(s string) {
	nb.Lock()
	defer nb.Unlock()
	fmt.Printf("%v -> %v\n", s, nb.n)
}

func (number *Number) increase() {
	fmt.Println("increase ", number.n)
	number.n++
	time.Sleep(time.Second * 1)
	fmt.Println("increase done ", number.n)
}
