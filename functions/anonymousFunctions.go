package functions

import "fmt"

func Calc()  {
	var numX int = 32
	var numY int = 243

	add := func(n1 int, n2 int) int {
		return n1 + n2 + numX + numY
	}

	fmt.Println(add(10, 25))
}
