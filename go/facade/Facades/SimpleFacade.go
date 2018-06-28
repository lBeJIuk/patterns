package Facades

import (
	"../Details"
)

type ActionsFacade struct {
	Detail1 Details.ConcreteDetail1
	Detail2 Details.ConcreteDetail2
}

func (facade ActionsFacade) MakeGeneralizedAktion(){
	facade.Detail1.MakeAction1_1()
	facade.Detail1.MakeAction1_2()

	facade.Detail2.MakeAction2_1()
	facade.Detail2.MakeAction2_2()
}