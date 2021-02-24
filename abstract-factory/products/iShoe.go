package products

//IShoe ..
type IShoe interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

//Shoe ..
type Shoe struct {
	Logo string
	Size int
}

//SetLogo ..
func (s *Shoe) SetLogo(logo string) {
	s.Logo = logo
}

//GetLogo ..
func (s *Shoe) GetLogo() string {
	return s.Logo
}

//SetSize ..
func (s *Shoe) SetSize(size int) {
	s.Size = size
}

//GetSize ..
func (s *Shoe) GetSize() int {
	return s.Size
}
