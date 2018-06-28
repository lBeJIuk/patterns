package Details

import "fmt"

type Detail1 interface {
	MakeAction1_1()
	MakeAction1_2()
}

type ConcreteDetail1 struct {
	Name string
}

func (detail ConcreteDetail1) MakeAction1_1() {
	fmt.Println("makeAction1_1")
}
func (detail ConcreteDetail1) MakeAction1_2() {
	fmt.Println("makeAction1_2")
}
