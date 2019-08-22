package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var nbAtomic = int32(0)

func main() {
	var nb = new(Number)

	go nb.safeIncrease()
	nb.printOut("1")
	go nb.safeIncrease()
	nb.printOut("2")
	go nb.safeIncrease()
	nb.printOut("3")
	go nb.safeIncrease()
	nb.printOut("4")
	go nb.safeIncrease()
	nb.printOut("5")

	time.Sleep(time.Second * 1)

	nb.printOut("FIN")

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

type Number struct {
	sync.RWMutex
	n int
}

func (number *Number) increase() {
	fmt.Println("increase ", number.n)
	number.n++
	time.Sleep(time.Second * 1)
	fmt.Println("increase done ", number.n)
}

func (number *Number) safeIncrease() {
	number.Lock()
	defer number.Unlock()
	number.increase()
}

func (number *Number) printOut(s string) {
	number.RLock()
	defer number.RUnlock()
	fmt.Printf("%v -> %v\n", s, number.n)
}

func increaseAtomic() {
	fmt.Println("increase ", atomic.LoadInt32(&nbAtomic))

	atomic.AddInt32(&nbAtomic, 1)

	time.Sleep(time.Second * 1)

	fmt.Println("increase done ", atomic.LoadInt32(&nbAtomic))
}

func printOutAtomic(s string) {
	fmt.Printf("%v -> %v\n", s, atomic.LoadInt32(&nbAtomic))
}
