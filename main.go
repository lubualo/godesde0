package main

import (
	"fmt"
	"github.com/lubualo/godesde0/variables"
)

func main() {
	state, text := variables.ConvertToText(1588)
	fmt.Println(state)
	fmt.Println(text)
}
