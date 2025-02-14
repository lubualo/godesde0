package models

import (
	"time"
)

type User struct {
	Id int
	Name string
	CreatedAtt time.Time
	Status bool
}

func (user *User) AddUser(id int, name string, date time.Time, status bool) {
	user.Id = id;
	user.Name = name;
	user.CreatedAtt = date;
	user.Status = status;
}

