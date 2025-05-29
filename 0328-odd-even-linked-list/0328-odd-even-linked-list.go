/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    evenHead := head.Next
    p1, p2 := head, evenHead

    for p2 != nil && p2.Next != nil {
        p1.Next = p2.Next
        p1 = p1.Next
        p2.Next = p1.Next
        p2 = p2.Next
    }
    p1.Next = evenHead
    return head
}