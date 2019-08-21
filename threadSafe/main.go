package main

import (
	"fmt"
	"sync"
	"time"
)

type Number struct {
	sync.Mutex
	n int
}

var nb = new(Number)

func main() {
	fmt.Println("0 -> ", nb.n)

	go increase()
	fmt.Println("1 -> ", nb.n)
	go increase()
	fmt.Println("2 -> ", nb.n)
	go increase()
	fmt.Println("3 -> ", nb.n)
	go increase()
	fmt.Println("4 -> ", nb.n)
	go increase()
	fmt.Println("5 -> ", nb.n)

	time.Sleep(time.Second * 1)

	fmt.Println("FIN -> ", nb.n)
}

func increase() {
	fmt.Println("increase ", nb.n)
	nb.n++
	time.Sleep(time.Second * 1)
	fmt.Println("increase done ", nb.n)
}
