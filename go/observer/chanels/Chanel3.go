package chanels

import "fmt"

type Chanel3 struct {
	Name string
}

func (chanel Chanel3) Update(){
	fmt.Println("News on ", chanel.Name)
}