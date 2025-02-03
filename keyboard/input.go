package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var legend string
var err error

func NunberProduct()  {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Write number 1: ")
	if scanner.Scan() {
		num1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("The input is incorrect " + err.Error())
		}
	}

	fmt.Println("Write number 2: ")
	if scanner.Scan() {
		num2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("The input is incorrect " + err.Error())
		}
	}

	fmt.Println("Write a legend: ")
	if scanner.Scan() {
		legend = scanner.Text()
	}

	fmt.Println(legend, num1 * num2)
}
