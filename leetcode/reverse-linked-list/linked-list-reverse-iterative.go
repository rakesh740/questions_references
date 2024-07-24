package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	node := head

	var newHead *ListNode

	for node != nil {

		curr := ListNode{Val: node.Val, Next: nil}
		curr.Next = newHead
		newHead = &curr

		node = node.Next

	}

	return newHead
}
