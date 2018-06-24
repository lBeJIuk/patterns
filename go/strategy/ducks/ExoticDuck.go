package ducks

import (
	"../flyStrategy"
	"../quackStrategy"
)

type ExoticDuck struct {
	BaseDuck
	flyStrategy.IFlyable
	quackStrategy.IQuackable
}
