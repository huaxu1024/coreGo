package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerId int
}

var dilbert Employee


type tree struct {
	value int
	left, right *tree
}

// tree start
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, v int) *tree {
	if t == nil {
		t = new(tree)
		t.value = v
		return t
	}
	if v < t.left.value {
		t.left = add(t.left, v)
	} else {
		t.right = add(t.right, v)
	}
	return t
}
// tree end


type Point struct {
	x, y int
}

func main() {
	p := Point{1, 2}
	fmt.Print(p)
	dilbert.Salary -= 2000
	position := &dilbert.Position
	*position = "Senior " + *position
}
