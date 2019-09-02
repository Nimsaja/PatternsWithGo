package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var count = 0

func main() {
	// something I found in the web:
	err := filepath.Walk("../.", walkerFunc)
	if err != nil {
		log.Println(err)
	}

	// my own stuff:
	err = filepath.Walk("../.", countMains)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("You've %v main.go under your folder\n", count)
}

func walkerFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(path, info.Size())
	return nil
}

func countMains(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if strings.Contains(path, "/main.go") {
		count++
	}
	return nil
}
