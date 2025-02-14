package arrays_slices

import "fmt"

var table2 []int = []int{2, 5}
var array [10]int = [10]int{6, 98, 21, 5555, 4, 7, -1, -5}

func ShowSlice() {
	fmt.Println(table2)
	slice1 := array[3:]
	fmt.Println(slice1)
	slice2 := array[:1]
	fmt.Println(slice2)
	slice3 := array[4:7]
	fmt.Println(slice3)
}

func Capacity()  {
	elements := make([]int, 5, 20)
	fmt.Printf("Length %d, capacity %d\n", len(elements), cap(elements))
	fmt.Println()

	nums := make([]int, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
		fmt.Printf("Length %d, capacity %d\n", len(nums), cap(nums))
	}
}
