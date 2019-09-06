package main

import (
	"fmt"
	"strings"
)

const (
	UPPER = 1  // upper case     0000 0001
	LOWER = 2  // lower case     0000 0010
	CAP   = 4  // capitalizes 	 0000 0100
	REV   = 8  // reverses       0000 1000
	TRIM  = 16 // no spaces      0001 0000
)

// & copies a bit to the result if it exists in both operands
// | copies a bit to the result if it exists in one of the operands

func main() {
	fmt.Println(procstr("HELLO PEOPLE!", LOWER|REV|CAP))        // 0000 1110
	fmt.Println(procstr("Hey there! How are you?", UPPER|TRIM)) // 0001 0001
}

func procstr(str string, conf byte) string {
	// query config bits
	if (conf & UPPER) != 0 {
		str = strings.ToUpper(str)
	}
	if (conf & LOWER) != 0 {
		str = strings.ToLower(str)
	}
	if (conf & CAP) != 0 {
		str = strings.Title(str)
	}
	if (conf & REV) != 0 {
		str = rev(str)
	}
	if (conf & TRIM) != 0 {
		str = strings.Replace(str, " ", "", -1)
	}
	return str
}

func rev(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}
