package main

import (
	e "github.com/lubualo/godesde0/excercise_interface"
	"github.com/lubualo/godesde0/models"
)

func main() {
	pedro := new(models.Man)
	e.HumansBreathing(pedro)
	martha := new(models.Woman)
	e.HumansBreathing(martha)
}
