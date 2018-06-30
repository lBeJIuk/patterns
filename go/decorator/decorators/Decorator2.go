package decorators

import (
	"../base"
	"fmt"
)

type Decorator2 struct {
	Base base.BaseInterface
}

func (decorator Decorator2)  MakeAction(){
	decorator.Base.MakeAction()
	fmt.Println("Decorator 2")
}