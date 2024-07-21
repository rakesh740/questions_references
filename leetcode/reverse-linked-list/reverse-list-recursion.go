/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList1(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return p
}

func reverseListIter(head *ListNode) *ListNode {

	var prev, curr *ListNode = nil, head

	// 5 -> 4 -> 3
	for prev.Next != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newhead, last := reverse(head)
	last.Next = nil

	return newhead
}

func reverse(head *ListNode) (*ListNode, *ListNode) {

	if head.Next == nil {
		return head, head
	}

	newhead, last := reverse(head.Next)
	last.Next = head
	last = last.Next

	return newhead, last
}
