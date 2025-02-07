package excercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Excercise02() string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	fmt.Println("Write a number:")
	if scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input, Let's try again!")
			return Excercise02()
		}
		for i := 0; i < 11; i++ {
			text += fmt.Sprintf("%d x %d = %d\n", num, i, num*i)
		}
	}
	return text
}
