package models

type Woman struct {
	Man
	IsMother bool
}

func (w *Woman) GetSex() string { return "Female" }
