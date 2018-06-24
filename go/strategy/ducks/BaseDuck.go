package ducks

import (
	"fmt"
)

type BaseDuck struct {
	Name string
}

func (duck BaseDuck) Display() {
	fmt.Println("display for: ", duck.Name)
}
