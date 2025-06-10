/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    if len(lists) == 1 {
        return lists[0]
    }

    l := lists[0]
    for i := 1; i < len(lists); i++ {
        l = mergeTwoLists(l, lists[i])
    }
    return l
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    newHead := &ListNode{Val: -999}
    p1, p2, curr := list1, list2, newHead
    
    for p1 != nil && p2 != nil {
        if p1.Val < p2.Val {
            curr.Next = p1
            p1 = p1.Next
        } else {
            curr.Next = p2
            p2 = p2.Next
        }
        curr = curr.Next
    }

    if p1 != nil {
        curr.Next = p1
    }
    if p2 != nil {
        curr.Next = p2
    }
    return newHead.Next
}