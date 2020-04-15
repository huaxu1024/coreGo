package reflectDiy

import (
	"fmt"
	"reflect"
)
type people interface {
	move()
}
type student struct {
	name string
}

func (*student) move() {}


func reflectDemo() {
	var stu people = &student{name: "huaxu"}
	t := reflect.TypeOf(stu)
	fmt.Println(t.String())
	fmt.Println(t)

	t = reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	of := reflect.ValueOf(3)
	fmt.Println(of)
	fmt.Printf("%v\n", of)
	fmt.Println(of.String())

	x := of.Interface()
	i := x.(int)
	fmt.Printf("%d\n", i)

	setX := 1
	rSetX := reflect.ValueOf(&setX).Elem()
	rSetX.SetInt(2)
	rSetX.Set(reflect.ValueOf(3))

	var setY interface{}
	rSetY := reflect.ValueOf(&setY).Elem()

	//rSetY.SetInt(2) // 崩溃， 在指向接口的Value上调用SetInt
	rSetY.Set(reflect.ValueOf(3))

	//rSetY.SetString("hello") // 崩溃， 在指向接口的Value上调用SetString
	rSetY.Set(reflect.ValueOf("hello"))
}
