package item

import (
	"fmt"

	"github.com/Kungfucoding23/design-paterns-go/observer/observer"
)

//Item struct
type Item struct {
	ObserverList []observer.Observer
	Name         string
	InStock      bool
}

//NewItem creates a new item
func NewItem(name string) *Item {
	return &Item{
		Name: name,
	}
}

//UpdateAvailability ..
func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.Name)
	i.InStock = true
	i.Notify()
}

//Register ...
func (i *Item) Register(obs observer.Observer) {
	i.ObserverList = append(i.ObserverList, obs)
}

//DeRegister ...
func (i *Item) DeRegister(obs observer.Observer) {
	i.ObserverList = removeFromSlice(i.ObserverList, obs)
}

//Notify ...
func (i *Item) Notify() {
	for _, observer := range i.ObserverList {
		observer.Update(i.Name)
	}
}

func removeFromSlice(observerList []observer.Observer, observerToRemove observer.Observer) []observer.Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
