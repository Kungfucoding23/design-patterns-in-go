package main

import (
	"github.com/Kungfucoding23/design-paterns-go/observer/customer"
	"github.com/Kungfucoding23/design-paterns-go/observer/item"
)

func main() {
	shirtItem := item.NewItem("Cool shirt")

	observer1 := &customer.Customer{ID: "register@gmail.com"}
	observer2 := &customer.Customer{ID: "register2@gmail.com"}

	shirtItem.Register(observer1)
	shirtItem.Register(observer2)

	shirtItem.UpdateAvailability()
}
