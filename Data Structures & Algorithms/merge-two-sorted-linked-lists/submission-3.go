/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	curNode := head

	for list1!=nil || list2!=nil {
		if list1 == nil {
			curNode.Next = list2
			if list2!=nil {
				list2 = list2.Next
			}
		} else if list2 == nil {
			curNode.Next = list1
			if list1!=nil {
				list1 = list1.Next
			}
		} else {
			if list1.Val < list2.Val {
				curNode.Next = list1
				if list1!=nil {
					list1 = list1.Next
				}
			} else {
				curNode.Next = list2
				if list2!=nil{
					list2 = list2.Next
				}
			}
		}
		curNode = curNode.Next
	}

	return head.Next
}
