package main

import "fmt"

// Node ...
type Node struct {
	val  int
	next *Node
}

//MyLinkedList ...
type MyLinkedList struct {
	head *Node
}

//Constructor Initialize your data structure here. */
func Constructor() MyLinkedList {

	return MyLinkedList{}
}

// Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (list *MyLinkedList) Get(index int) int {

	if index < 0 {
		return -1
	}
	var node *Node
	node = list.head

	if index == 0 {
		return node.val
	}

	for i := 1; i <= index; i++ {
		node = node.next

		if node == nil {
			return -1
		}
	}

	return node.val
}

//AddAtHead Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (list *MyLinkedList) AddAtHead(val int) {

	node := Node{val: val, next: nil}

	node.next = list.head
	list.head = &node
}

//AddAtTail Append a node of value val to the last element of the linked list. */
func (list *MyLinkedList) AddAtTail(val int) {
	node := list.head

	newNode := Node{val: val, next: nil}

	if node == nil {
		list.head = &newNode
		return
	}

	for node.next != nil {
		node = node.next
	}

	node.next = &newNode
}

//AddAtIndex Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (list *MyLinkedList) AddAtIndex(index int, val int) {

	var node *Node
	newNode := Node{val: val, next: nil}

	node = list.head

	if index == 0 {
		list.AddAtHead(val)
		return
	}

	for i := 1; i < index; i++ {
		node = node.next

		if node == nil {
			return
		}
	}

	newNode.next = node.next
	node.next = &newNode
}

//DeleteAtIndex Delete the index-th node in the linked list, if the index is valid. */
func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}

	if index == 0 {
		list.head = list.head.next
		return
	}

	var node *Node

	for i := 0; i < index; i++ {

		if i == 0 {
			node = list.head
		} else {

			node = node.next

		}

		if node.next == nil {
			return
		}

	}

	node.next = node.next.next
}

//Print ...
func (list *MyLinkedList) Print() {

	node := list.head

	for node != nil {

		fmt.Printf("%+v ->", node.val)

		node = node.next

	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {

	link := Constructor()

	link.AddAtHead(9)

	link.Print()
	fmt.Print("\n")

	link.Get(4)

	link.AddAtHead(5)
	link.AddAtHead(5)
	link.AddAtHead(5)
	link.AddAtHead(5)
	link.AddAtHead(5)

	link.AddAtTail(10)
	link.Get(5)

	link.DeleteAtIndex(6)

	link.DeleteAtIndex(4)

	link.Print()
	fmt.Print("\n")

	// link.AddAtIndex(1, 78)

	// link.Print()

	// fmt.Print("\n")
}

/**
 */
