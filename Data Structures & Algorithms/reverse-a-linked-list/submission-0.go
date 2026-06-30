/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prevNode, nextNode *ListNode
	curNode := head
	nextNode = curNode.Next
	
	for curNode != nil {
		// fmt.Println(curNode.Val)
		curNode.Next = prevNode
		prevNode = curNode
		curNode = nextNode
		if curNode != nil {
			nextNode = curNode.Next
		}
	}

	return prevNode
}
