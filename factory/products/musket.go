package products

//Musket struct
type Musket struct {
	Gun
}

//NewMusket crea una nueva musket
func NewMusket() IGun {
	return &Musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
