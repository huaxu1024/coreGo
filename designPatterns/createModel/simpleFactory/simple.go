package simpleFactory

import "fmt"

/**
	《简单工厂》
	go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
	NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。
	在这个simpleFactory包中只有API 接口和NewAPI函数为包外可见，封装了实现细节。
 */

type API interface {
	Say(name string) string
}

func NewAPI(i int) API{
	if i == 1 {
		return &HiApi{}
	} else if i == 2 {
		return &HelloApi{}
	}
	return nil
}

type HiApi struct {
}

func (*HiApi) Say(name string) string {
	return fmt.Sprintf("Hi, " + name)
}

type HelloApi struct {
}

func (*HelloApi) Say(name string) string {
	return fmt.Sprintf("Hello, " + name)
}


