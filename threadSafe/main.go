package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Number struct {
	sync.RWMutex
	n int
}

var nb = new(Number)

var nbAtomic = int32(0)

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

	fmt.Println("***********************************")

	go increaseAtomic()
	printOutAtomic("1")
	go increaseAtomic()
	printOutAtomic("2")
	go increaseAtomic()
	printOutAtomic("3")
	go increaseAtomic()
	printOutAtomic("4")
	go increaseAtomic()
	printOutAtomic("5")

	time.Sleep(time.Second * 1)

	printOutAtomic("FIN")
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

func increaseAtomic() {
	nn := atomic.LoadInt32(&nbAtomic)
	fmt.Println("increase ", nn)

	nn++
	atomic.StoreInt32(&nbAtomic, nn)
	time.Sleep(time.Second * 1)

	fmt.Println("increase done ", nn)
}

func printOutAtomic(s string) {
	fmt.Printf("%v -> %v\n", s, atomic.LoadInt32(&nbAtomic))
}
