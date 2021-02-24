package products

//IShirt ..
type IShirt interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

//Shirt ..
type Shirt struct {
	Logo string
	Size int
}

//SetLogo ..
func (s *Shirt) SetLogo(logo string) {
	s.Logo = logo
}

//GetLogo ...
func (s *Shirt) GetLogo() string {
	return s.Logo
}

//SetSize ..
func (s *Shirt) SetSize(size int) {
	s.Size = size
}

//GetSize ...
func (s *Shirt) GetSize() int {
	return s.Size
}
