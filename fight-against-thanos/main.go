package main

import (
	"fmt"
)

func main() {
	army1 := nationalArmy{"Air Forces", 10, 1000}
	army2 := nationalArmy{"Marines", 20, 1000}
	army3 := nationalArmy{"Space Forces", 40, 1000}
	army4 := heroArmy{"Iron Man", 100000, 50000}
	army5 := heroArmy{"Thor", 200000, 0}
	// chuck := chuckArmy{}

	armies := []army{army1, army2, army3, army4, army5}
	fightAgainstThanos(armies)
}

type army interface {
	name() string
	power() int
}

type nationalArmy struct {
	armyName  string
	unitPower int
	soldiers  int
}

func (n nationalArmy) name() string {
	return n.armyName
}

func (n nationalArmy) power() int {
	return n.unitPower * n.soldiers
}

type heroArmy struct {
	heroName       string
	heroPower      int
	sideKicksPower int
}

func (h heroArmy) name() string {
	return fmt.Sprintf("Army with %s", h.heroName)
}

func (h heroArmy) power() int {
	return h.heroPower + h.sideKicksPower
}

type chuckArmy struct {
}

func (c chuckArmy) name() string {
	return "The greatest Chuck Norris"
}

func (c chuckArmy) power() int {
	return 9999999
}

func fightAgainstThanos(armies []army) {
	thanosLife := 9999999
	for _, army := range armies {
		fmt.Printf("%s fight against Thanos with a force of %d\n", army.name(), army.power())
		thanosLife -= army.power()
	}
	if thanosLife <= 0 {
		fmt.Println("The terrible Thanos has been defeated")
	} else {
		fmt.Printf("Thanos destroy entire universe, and his life is %d yet\n", thanosLife)
	}
}
