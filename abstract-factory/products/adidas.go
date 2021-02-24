package products

//Adidas struct
type Adidas struct {
}

//MakeShoe ..
func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			Logo: "adidas",
			Size: 39,
		},
	}
}

//MakeShirt ..
func (a *Adidas) MakeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}
