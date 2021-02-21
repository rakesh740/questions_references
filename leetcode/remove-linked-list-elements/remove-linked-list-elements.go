/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeElements(head *ListNode, val int) *ListNode {
    
      
    if head == nil{
        return nil
    }
    
   var newHead *ListNode
    
    curr := head.Next
  
    
    if head.Val != val {
        newHead = &ListNode{Val: head.Val, Next: nil}
    }
    secondPointer := newHead
    
    for curr != nil {
        if curr.Val != val { 
             node := &ListNode{Val: curr.Val, Next: nil}
            if newHead == nil {
                secondPointer = node
                newHead = secondPointer
                
            } else {

                secondPointer.Next = node
                secondPointer = secondPointer.Next
             }
            }
       
        curr = curr.Next
    } 
    
    return newHead
    
}
