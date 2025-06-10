/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    n := 0
    for p := head; p != nil; p = p.Next {
        n++
    }

    dummyHead := &ListNode{Next: head}
    node0, curr := dummyHead, head
    var pre *ListNode

    for ; n >= k; n -= k {
        pre = nil

        for i := 0; i < k; i++ {
            next := curr.Next
            curr.Next = pre
            pre = curr
            curr = next
        }

        next := node0.Next
        next.Next = curr
        node0.Next = pre

        node0 = next
    }

    return dummyHead.Next
}