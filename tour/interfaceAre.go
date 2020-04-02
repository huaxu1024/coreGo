package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}
// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t T) M() {
	fmt.Printf("this is T %s", t.S)
}


type Z struct {
	X string
}

func (t Z) M() {
	fmt.Printf("this is Z %s", t.X)
}

func goRUn() {
	var i I = Z{"hello"}
	i.M()
}
