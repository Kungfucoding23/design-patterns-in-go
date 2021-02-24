package main

import (
	"github.com/Kungfucoding23/design-paterns-go/factory/products"
)

func main() {
	ak47, _ := products.GetGun("ak47")
	musket, _ := products.GetGun("musket")

	products.PrintDetails(ak47)
	products.PrintDetails(musket)
}
