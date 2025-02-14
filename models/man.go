package models

type Man struct {
	Age       int
	Height    float32
	Weight    float32
	Thinking  bool
	Breathing bool
	Eating    bool
	Alive   bool
}

func (m *Man) IsAlive() bool { return m.Alive }
func (m *Man) Breath() { m.Breathing = true }
func (m *Man) Think()  { m.Thinking = true }
func (m *Man) Eat()    { m.Eating = true }
func (m *Man) GetSex() string { return "Male" }
