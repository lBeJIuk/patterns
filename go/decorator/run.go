package main

import (
	"./base"
	"./decorators"
	"fmt"
)

func main() {
	var base1 base.BaseInterface

	base1 = base.Base{"Base1"}
	base1.MakeAction()
	fmt.Println("======")

	base1 = decorators.Decorator1{base1}
	base1.MakeAction()
	fmt.Println("======")

	base1 = decorators.Decorator2{base1}
	base1.MakeAction()
	fmt.Println("======")
}

