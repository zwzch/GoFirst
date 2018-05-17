package Test

import "fmt"

type Signal interface {
	String() string
	Signal()
}

func Test(){
	fmt.Println("heihei")
}
