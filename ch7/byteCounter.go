package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func test() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

type Stringer interface {
	String() string
}

func coryCore() {
	a := []int{1,2,3,4}
	b := make([]int, 1)
	copy(b, a)
	fmt.Println(a)
	fmt.Printf("%v", b)
}