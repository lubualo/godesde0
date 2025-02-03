package main

import (
	"fmt"
	"github.com/lubualo/godesde0/excercises"
)

func main() {
	number, text := excercises.Excercise01("12345")
	fmt.Println(number)
	fmt.Println(text)

	number, text = excercises.Excercise01("123A45")
	fmt.Println(number)
	fmt.Println(text)
}
