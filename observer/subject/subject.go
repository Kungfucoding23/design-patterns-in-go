package subject

import "github.com/Kungfucoding23/design-paterns-go/observer/observer"

//Subject interface
type Subject interface {
	Register(observer observer.Observer)
	DeRegister(observer observer.Observer)
	Notify()
}
