/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummyHead := &ListNode{Val:-1}
    dummyHead.Next = head 
    p1 := dummyHead
    for i := 0; i < left-1; i++ {
        p1 = p1.Next
    }
    p2 := p1.Next
    q1 := p1
    for i := 0; i < right - left + 1; i++ {
        q1 = q1.Next
    }
    q2 := q1.Next

    p1.Next = nil
    q1.Next = nil

    reverse(p2)

    p1.Next = q1
    p2.Next = q2


    return dummyHead.Next
}

func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    newHead := reverse(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}