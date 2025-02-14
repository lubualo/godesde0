package excercise_interface

import (
	"fmt"
	"github.com/lubualo/godesde0/interfaces"
)

func HumansBreathing(hu interfaces.Human) {
	hu.Breath()
	fmt.Printf("I'm a %s and I'm breathing\n", hu.GetSex())
}

func AliveLiferforms(lf interfaces.Lifeform) bool {
	return lf.IsAlive()
}
