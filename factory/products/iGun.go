package products

//IGun define todos los métodos con los que debería contar un arma
type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

//Gun implementa la interfaz iGun
type Gun struct {
	name  string
	power int
}

//SetName ...
func (g Gun) SetName(name string) {
	g.name = name
}

//GetName ...
func (g Gun) GetName() string {
	return g.name
}

//SetPower ..
func (g Gun) SetPower(power int) {
	g.power = power
}

//GetPower ...
func (g Gun) GetPower() int {
	return g.power
}
