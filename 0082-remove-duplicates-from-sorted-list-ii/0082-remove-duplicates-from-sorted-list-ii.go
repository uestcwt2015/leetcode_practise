/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    dummyHead := &ListNode{Next: head, Val: 999}
    p1 := dummyHead
    for ; p1.Next != nil && p1.Next.Next != nil; {
        if p1.Next.Val == p1.Next.Next.Val {
            x := p1.Next.Val
            for p1.Next != nil && p1.Next.Val == x {
                p1.Next = p1.Next.Next
            }
        } else {
            p1 = p1.Next
        }
    }
    return dummyHead.Next
}