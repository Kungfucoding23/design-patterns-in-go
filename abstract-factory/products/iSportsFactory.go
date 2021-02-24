package products

import "fmt"

//ISportsFactory interface
type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}

//GetSportsFactory retorna el producto segun la marca
func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

//PrintShoeDetails ...
func PrintShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

//PrintShirtDetails ...
func PrintShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
