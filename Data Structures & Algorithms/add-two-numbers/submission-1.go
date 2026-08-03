/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var additionBuff, sum int
	
	result := &ListNode{
		Val: (l1.Val+l2.Val)%10,
		Next: nil,
	}
	resultPtr := result
	additionBuff = (l1.Val+l2.Val)/10
	l1 = l1.Next
	l2 = l2.Next

	// fmt.Println(l1, l2)
	for l1!=nil || l2!=nil {
		if l1==nil {
			sum = l2.Val+additionBuff
			l2 = l2.Next
		} else if l2==nil {
			sum = l1.Val+additionBuff
			l1 = l1.Next
		} else {
			sum = l1.Val+l2.Val+additionBuff
			l1 = l1.Next
			l2 = l2.Next
		}
		// fmt.Println(l1, l2, additionBuff, sum)

		resultPtr.Next = &ListNode{
			Val: sum%10,
			Next: nil,
		}
		additionBuff = sum/10
		resultPtr = resultPtr.Next
	}

	if additionBuff!=0 {
		resultPtr.Next = &ListNode{
			Val: additionBuff,
			Next: nil,
		}
	}

	return result
}
