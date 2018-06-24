package ducks

import (
	"../flyStrategy"
	"../quackStrategy"
)

type UpgradableDuck struct {
	BaseDuck
	IFlyable flyStrategy.IFlyable
	IQuackable quackStrategy.IQuackable
}

func (upgradableDuck UpgradableDuck) Quack(){
	upgradableDuck.IQuackable.Quack()
}

func (upgradableDuck UpgradableDuck) Fly(){
	upgradableDuck.IFlyable.Fly()
}
