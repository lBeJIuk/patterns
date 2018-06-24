package main

import (
	"./ducks"
	"./flyStrategy"
	"./quackStrategy"
	"fmt"
)

func main() {

	simpleDuck := ducks.SimpleDuck{ducks.BaseDuck{"Test duck 1"}, flyStrategy.SimpleFly{}, quackStrategy.SimpleQuack{}}
	exoticDuck := ducks.ExoticDuck{ducks.BaseDuck{"Test duck 2"}, flyStrategy.ExoticFly{}, quackStrategy.ExoticQuack{}}

	upgradable := ducks.UpgradableDuck{ducks.BaseDuck{"Test duck 3"}, flyStrategy.NoFly{}, quackStrategy.NoQuack{}}

	ducksList := make([]ducks.IDuck, 0)
	ducksList = append(ducksList, simpleDuck)
	ducksList = append(ducksList, exoticDuck)
	ducksList = append(ducksList, upgradable)

	for i := 0; i <= len(ducksList) - 1; i++ {
		ducksList[i].Display()
		ducksList[i].Quack()
		ducksList[i].Fly()
		fmt.Println("===")
	}
// @TODO SetFlyStrategy???

}
