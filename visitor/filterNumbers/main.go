package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// test
	fmt.Println("even array: ", convert(a, evenFilter))
	fmt.Println("uneven array: ", convert(a, unevenFilter))
}

type filter func(int) bool

func convert(a []int, f filter) (l []int) {
	for _, n := range a {
		if f(n) {
			l = append(l, n)
		}
	}
	return l
}

func evenFilter(n int) bool { return n%2 == 0 }

func unevenFilter(n int) bool { return n%2 != 0 }
