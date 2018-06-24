package ducks

import (
	"../flyStrategy"
	"../quackStrategy"
)

type UpgradableDuck struct {
	BaseDuck
	flyStrategy.IFlyable
	quackStrategy.IQuackable
}

//func (upgradableDuck *UpgradableDuck) SetFlyStrategy(flyStrategy flyStrategy.IFlyable) {
//	upgradableDuck.Fly = flyStrategy.Fly
//}
