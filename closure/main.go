package main

import "fmt"

func main() {
	f := once()
	f()
	f()

	// //wrong! calls once another two times
	// once()()
	// once()()

	// g := once()
	// g()
	// g()

	f = once()
	f()
	f()

	finc := inc()
	finc()
	finc()

}

func once() func() {
	executed := false
	return func() {
		if !executed {
			fmt.Println("call function")
			executed = true
		}
	}
}

func inc() func() {
	i := 0
	return func() {
		fmt.Println(i)
		i++
	}
}
