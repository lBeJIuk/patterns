package main

import (
	"./Details"
	"./Facades"
	"fmt"
)



func main() {
	d1 := Details.ConcreteDetail1{"d1"}
	d2 := Details.ConcreteDetail2{"d2"}

	d1.MakeAction1_1()
	d1.MakeAction1_2()
	d2.MakeAction2_1()
	d2.MakeAction2_2()

	// OR with facade
	fmt.Println("=======")
	facade := Facades.ActionsFacade{Details.ConcreteDetail1{"d1"}, Details.ConcreteDetail2{"d2"}}
	facade.MakeGeneralizedAktion()

}