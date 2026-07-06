/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func revertList(head *ListNode) *ListNode {
	var prev, next *ListNode
	curr := head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// revert the list
	head = revertList(head)
	// fmt.Println(head.Val)

	currNode := head
	for i:=1; i<n-1; i++ {
		currNode = currNode.Next
	}
	if n==1 && currNode==head {
		if currNode.Next==nil {
			return nil
		}

		newHead := currNode.Next
		currNode.Next = nil
		head = newHead
	} else {
		currNode.Next = currNode.Next.Next
	}

	// revert the list back
	head = revertList(head)
	return head
}
