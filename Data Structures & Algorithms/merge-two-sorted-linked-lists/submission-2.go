/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1==nil && list2==nil {
		return nil
	}

	var newHead *ListNode
	if list1==nil {
		newHead = list2
		list2 = list2.Next
	} else if list2==nil {
		newHead = list1
		list1 = list1.Next
	} else {
		if list1.Val <= list2.Val {
			newHead = list1
			list1 = list1.Next
		} else {
			newHead = list2
			list2 = list2.Next
		}
	}
	sortedList := newHead

	for list1 != nil || list2 != nil {
		if list1==nil {
			// fmt.Println("List1 is nil")
			sortedList.Next = list2
			list2 = list2.Next
		} else if list2==nil {
			// fmt.Println("List2 is nil")
			sortedList.Next = list1
			list1 = list1.Next
			// sortedList = list1
		} else if list1.Val < list2.Val {
			// fmt.Println("List1.Val < List2.Val")
			sortedList.Next = list1
			list1 = list1.Next
			// sortedList = list1
		} else {
			// fmt.Println("List1.Val > List2.Val")
			sortedList.Next = list2
			list2 = list2.Next
			// sortedList = list2
		}
		sortedList = sortedList.Next
	}

	return newHead
}
