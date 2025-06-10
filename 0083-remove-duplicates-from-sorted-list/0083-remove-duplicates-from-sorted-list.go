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
    curr := head
    p := curr.Next
    for ; p != nil; p = p.Next {
        if p.Val == curr.Val {
            continue
        }
        curr.Next = p
        curr = curr.Next
    }
    curr.Next = nil
    return head
}