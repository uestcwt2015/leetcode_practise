func hasCycle(head *ListNode) bool {
    p1, p2 := head, head

    if head == nil || head.Next == nil {
        return false
    }

    for p2 != nil && p2.Next != nil {
        p1 = p1.Next
        p2 = p2.Next.Next
        if p1 == p2 {
            return true
        }
    }

    return false
}