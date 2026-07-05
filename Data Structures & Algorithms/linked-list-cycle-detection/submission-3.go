/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head==nil {
		return false
	}

	fastPtr, slowPtr := head, head
	for fastPtr.Next!=nil && fastPtr.Next.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if fastPtr==slowPtr {
			return true
		}
	}

	return false
}
