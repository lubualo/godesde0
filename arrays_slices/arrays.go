package arrays_slices

import "fmt"

var table [10]int = [10]int{0, 1, 2, 3, 4}
var matrix [20][30]int

func ShowArrays()  {
	table[7] = 33
	table[2] = 54

	table2 := [10]string{"david", "dani", "lucas", "trini"}
	fmt.Println(table)
	fmt.Println()
	fmt.Println(table2)
	fmt.Println()

	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}
	fmt.Println()

	matrix[7][24] = 15
	fmt.Println(matrix)
	fmt.Println()
}
