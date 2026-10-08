/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var prevNode, nextNode *ListNode
	curNode := head

	for curNode != nil {
		nextNode = curNode.Next
		curNode.Next = prevNode
		prevNode = curNode
		curNode = nextNode
		
		// if prevNode==nil {
		// 	curNode
		// }
	}

	return prevNode
}
