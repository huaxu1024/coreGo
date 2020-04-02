package main

import (
	"fmt"
	"image/color"
)

type ListNode struct {
	Value int
	Next *ListNode
}

func (list *ListNode) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Next.Sum()
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func test2() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)

	cp.Point.Y = 2
	fmt.Println(cp.Y)
}