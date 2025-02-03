package excercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Excercise02() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Write a number:")
	if scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		fmt.Println()
		if err != nil {
			fmt.Println("Invalid input, Let's try again!")
			Excercise02()
			return
		}
		for i := 0; i < 11; i++ {
			fmt.Printf("%d x %d = %d\n", num, i, num*i)
		}
	}

}
