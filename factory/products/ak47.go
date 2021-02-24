package products

//Ak47 struct
type Ak47 struct {
	Gun
}

//NewAk47 crea una ak47
func NewAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
