package main

import "fmt"

func main() {
	handleError(doSomething(0))
	handleError(doSomething(-1))
	handleError(doSomething(1))
}

func doSomething(i int) error {
	switch {
	case i == 0:
		return fmt.Errorf("****0*****")
	case i > 0:
		return fmt.Errorf("****i > 0 **** %v", i)
	case i < 0:
		return fmt.Errorf("****i < 0 **** %v", i)
	}

	return fmt.Errorf("")
}

func handleError(err error) {
	fmt.Printf("Error: %v\n", err)
}
