package observer

//Observer interface
type Observer interface {
	Update(string)
	GetID() string
}
