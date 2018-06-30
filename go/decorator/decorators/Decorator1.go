package decorators

import (
	"../base"
	"fmt"
)

type Decorator1 struct {
	Base base.BaseInterface
}

func (decorator Decorator1)  MakeAction(){
	decorator.Base.MakeAction()
	fmt.Println("Decorator 1")
}