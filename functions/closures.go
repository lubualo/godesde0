package functions

import "fmt"

func table(value int) (func() int)  {
	number := value
	sequence := 0
	return func() int {
		sequence++
		return number * sequence
	}
}

func CallClosuere()  {
	tableOf := 2
	MyTable := table(tableOf)
	for i := 0; i < 10; i++ {
		fmt.Println(MyTable())
	}
} 
