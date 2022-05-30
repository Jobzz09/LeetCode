package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	start := head
	length := 1

	for start.Next != nil {
		length++
		start = start.Next
	}

	start = head
	length /= 2
	for length != 0 {
		start = start.Next
		length--
	}

	return start
}

func main() {
	var head *ListNode
	for i := 0; i < 10; i++ {
		var newNode ListNode
		newNode.Val = i
		newNode.Next = head

		head = &newNode
	}
	fmt.Println("Before:")
	start := head
	for start != nil {
		fmt.Print(start.Val, " ")
		start = start.Next
	}

	fmt.Println("After:")
	a := middleNode(head)
	start = a
	for start != nil {
		fmt.Print(start.Val, " ")
		start = start.Next
	}
}
