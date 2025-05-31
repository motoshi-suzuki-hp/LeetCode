func mergeNodes(head *ListNode) *ListNode {
    dummy := &ListNode{Val: 0}
    tail := dummy

    cur := head.Next
    sum := 0

    for cur != nil {
        if cur.Val != 0 {
            sum += cur.Val
        } else {
            if sum != 0 {
                newNode := &ListNode{Val: sum}
                tail.Next = newNode
                tail = newNode
            }
            sum = 0
        }
        cur = cur.Next
    }

    return dummy.Next
}