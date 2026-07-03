/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head==nil {
		return nil
	}

	curNode := head
	var prevNode, nextNode *ListNode
	for curNode!=nil {
		nextNode = curNode.Next
		curNode.Next = prevNode
		prevNode = curNode
		curNode = nextNode
	}
	return prevNode
}
