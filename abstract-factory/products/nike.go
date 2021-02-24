package products

//Nike ...
type Nike struct {
}

//MakeShoe ...
func (n *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			Logo: "nike",
			Size: 39,
		},
	}
}

//MakeShirt ..
func (n *Nike) MakeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
