package interfaces

type Animal interface {
	Lifeform
	Breath()
	Eat()
	IsCarnivore() bool
}