/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    
    if head == nil || head.Next == nil{ return head }
    p := reverseList(head.Next)
    
    head.Next.Next = head
    head.Next = nil
    return p;

}

func Print(head *ListNode) {

	node := head

	for node != nil {

		fmt.Printf("%+v ->", node.Val)

		node = node.Next

	}
}
