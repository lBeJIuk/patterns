package chanels

import "fmt"

type Chanel1 struct {
	Name string
}

func (chanel Chanel1) Update(){
	fmt.Println("News on ", chanel.Name)
}