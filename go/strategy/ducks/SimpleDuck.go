package ducks

import (
	"../flyStrategy"
	"../quackStrategy"
)

type SimpleDuck struct {
	BaseDuck
	flyStrategy.IFlyable
	quackStrategy.IQuackable
}

