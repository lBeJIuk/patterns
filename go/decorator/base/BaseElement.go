package base

import "fmt"

type Base struct {
	Name string
}

func  (base Base) MakeAction(){
	fmt.Println("Make aktion")
}

type BaseInterface interface {
	MakeAction()
}