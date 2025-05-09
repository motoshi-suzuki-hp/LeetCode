/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    current := head

    for current != nil && current.Next != nil {
        nextNode := current.Next

        gcdNode := &ListNode{
            Val:  gcd(current.Val, nextNode.Val),
            Next: nextNode,
        }

        current.Next = gcdNode

        current = nextNode
    }

    return head
}