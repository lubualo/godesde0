package iterations

import (
	"fmt"
)

func Iterate()  {
	for i := 100; i > 10; i -= 15 {
		if i == 55 {
			continue
		}
		fmt.Println(i)
	}
}
