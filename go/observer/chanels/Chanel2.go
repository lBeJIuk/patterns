package chanels

import "fmt"

type Chanel2 struct {
	Name string
}

func (chanel Chanel2) Update(){
	fmt.Println("News on ", chanel.Name)
}