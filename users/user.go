package users

import (
	"fmt"
	"time"
	"github.com/lubualo/godesde0/models"
)

func RegisterUser()  {
	u := new(models.User)
	u.AddUser(345445640, "Digimon", time.Now(), true)
	fmt.Println(u)
}
