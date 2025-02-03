package excercises

import (
	"strconv"
)

func Excercise01(text string) (int, string)  {
	number, error := strconv.Atoi(text)
	if error != nil {
		return 0, "Error parsing the number"
	}
	var response string
	if number > 100 {
		response = "Greater than 100"
	} else if number < 100 {
		response = "Smaller than 100"
	} else {
		response = "Equals to 100"
	}
	return number, response
}