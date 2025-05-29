/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return head
    }
    tail, length := head, 1
    for tail.Next != nil {
        length++
        tail = tail.Next
    }
    tail.Next = head

    for i := 0; i < length - k % length; i++ {
        head = head.Next
        tail = tail.Next
    }

    tail.Next = nil
    return head
}