/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    newHead := &ListNode{Val: -1}
    p, q := head, newHead
    for p != nil && p.Next != nil {
        a, b := p, p.Next
        p = b.Next
        q.Next = b
        b.Next = a
        q = a
    }

    if p != nil {
        q.Next = p
        q = p
    }
    q.Next = nil
    return newHead.Next
}