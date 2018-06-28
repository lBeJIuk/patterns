package Details

import "fmt"

type Detail2 interface {
	MakeAction2_1()
	MakeAction2_2()
}

type ConcreteDetail2 struct {
	Name string
}

func (detail ConcreteDetail2) MakeAction2_1() {
	fmt.Println("makeAction2_1")
}
func (detail ConcreteDetail2) MakeAction2_2() {
	fmt.Println("makeAction2_2")
}
