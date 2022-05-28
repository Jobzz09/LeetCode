package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head

	for i := 0; i < n+1; i++ {
		if fast == nil {
			return head.Next
		}
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return head
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
	a := removeNthFromEnd(head, 2)
	start = a
	for start != nil {
		fmt.Print(start.Val, " ")
		start = start.Next
	}
}
