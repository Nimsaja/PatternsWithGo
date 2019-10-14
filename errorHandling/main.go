package main

import "fmt"

type isZeroError struct {
}

func (e isZeroError) Error() string {
	return fmt.Sprintf("Error: is zero")
}

type isTooBigError struct {
	arg int
}

func (e isTooBigError) Error() string {
	return fmt.Sprintf("Error: is too big: %v", e.arg)
}

type isTooLowError struct {
	arg int
}

func (e isTooLowError) Error() string {
	return fmt.Sprintf("Error: is too low: %v", e.arg)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print("Please put in a number: ")
		var guess int
		_, err := fmt.Scanf("%d", &guess)

		if err != nil {
			handleError(err)
		} else {
			handleError(doSomething(guess))
		}
	}
}

func doSomething(i int) error {
	switch {
	case i == 0:
		return isZeroError{}
	case i > 0:
		return isTooBigError{i}
	case i < 0:
		return isTooLowError{i}
	default:
		return nil
	}
}

func handleError(err error) {
	switch err.(type) {
	case isZeroError:
		fmt.Printf("%v. Please choose a number !=0 !\n", err)
	case isTooBigError:
		fmt.Printf("%v. Please choose a number <= 0!\n", err)
	case isTooLowError:
		fmt.Printf("%v. Please choose a number >= 0!\n", err)
	default:
		fmt.Printf("Error: %v\n", err)
	}
}
