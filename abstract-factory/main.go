package main

import "github.com/Kungfucoding23/design-paterns-go/abstract-factory/products"

func main() {
	adidasFactory, _ := products.GetSportsFactory("adidas")
	nikeFactory, _ := products.GetSportsFactory("nike")
	// nikeFactory, _ := products.GetSportsFactory("tomato")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	products.PrintShoeDetails(nikeShoe)
	products.PrintShirtDetails(nikeShirt)
	products.PrintShoeDetails(adidasShoe)
	products.PrintShirtDetails(adidasShirt)
}
