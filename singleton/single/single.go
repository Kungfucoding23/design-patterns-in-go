package single

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

//Single struct
type Single struct {
}

//SingleInstance ...
var SingleInstance *Single

//GetInstance ..
func GetInstance() *Single {
	if SingleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if SingleInstance == nil {
			fmt.Println("Creating single instance now.")
			SingleInstance = &Single{}
		} else {
			fmt.Println("Single instance already created")
		}
	} else {
		fmt.Println("Single instance already created")
	}
	return SingleInstance
}
