/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    var lists []*ListNode

	for head!=nil{
		lists = append(lists, head)
		head = head.Next
	}

	left := min(1, len(lists)-1) 
	right := len(lists)-1
	curr := lists[0]
	head = lists[0]

	for left <= right {
		curr.Next = lists[right]
		curr.Next.Next = lists[left]
		curr = curr.Next.Next

		left += 1
		right -= 1
	}
	curr.Next = nil
}
