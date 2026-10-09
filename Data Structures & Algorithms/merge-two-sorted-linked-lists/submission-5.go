/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1==nil && list2==nil{
		return nil
	}
	
	var sortedList, currNode *ListNode

	if list1!=nil && (list2==nil || list1.Val<=list2.Val) {
		sortedList = list1
		currNode = list1
		list1 = list1.Next
	} else {
		sortedList = list2
		currNode = list2
		list2 = list2.Next
	}

	for list1!=nil || list2!=nil {
		if list1==nil {
			currNode.Next = list2
			list2 = list2.Next
		} else if list2==nil {
			currNode.Next = list1
			list1 = list1.Next
		} else {
			if list1.Val <= list2.Val {
				currNode.Next = list1
				list1 = list1.Next
			} else {
				currNode.Next = list2
				list2 = list2.Next
			}
		}
		currNode = currNode.Next
	}

	return sortedList
}
