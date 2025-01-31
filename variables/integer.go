package variables

import "fmt"

func ShowIntegers()  {
	var commonint int
	var commonint32 int32
	intof32 := int32(10)
	intof64 := int64(100)

	fmt.Println("commonint = ", commonint)
	fmt.Println("commonint32 = ", commonint32)
	fmt.Println("intof32 = ", intof32)
	fmt.Println("intof64 = ", intof64)
}