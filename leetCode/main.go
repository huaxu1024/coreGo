package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reserve(head *ListNode) *ListNode{
	var res, temp *ListNode
	for head != nil {
		temp = head.Next
		head.Next = res
		res = head
		head = temp
	}
	return res
}

func main() {
	node1 := new(ListNode)
	node1.Val = 1
	node2 := new(ListNode)
	node2.Val = 2
	node3 := new(ListNode)
	node3.Val = 3
	node4 := new(ListNode)
	node4.Val = 4

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = nil
	res := reserve(node1)

	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}