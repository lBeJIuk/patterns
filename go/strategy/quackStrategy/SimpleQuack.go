package quackStrategy

import "fmt"

type SimpleQuack struct {
}

func (simpleQuack SimpleQuack) Quack() {
	fmt.Println("Simple Quack")
}