/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    fastPt, slowPt := head, head

    for fastPt!=nil {
        slowPt = slowPt.Next
        fastPt = fastPt.Next
        if fastPt!=nil {
            fastPt = fastPt.Next
        }

        if slowPt!=nil && fastPt!=nil && slowPt==fastPt {
            return true
        }
    }

    return false
}
