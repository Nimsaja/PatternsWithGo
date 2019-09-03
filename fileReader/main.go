package main

import (
	"fmt"

	"github.com/Nimsaja/PatternsWithGo/fileReader/tools"
)

func main() {
	text, err := tools.Load("file.txt")

	if err == nil {
		fmt.Println(text)
	}
}
