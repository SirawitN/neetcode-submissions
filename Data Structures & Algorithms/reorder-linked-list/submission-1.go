/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head==nil {
        return
    }

    // Get the middle index of the list
    slow, fast := head, head
    for fast!=nil && fast.Next!=nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    mid := slow.Next
    // Reverse the second half
    curr := mid
    var prev, next *ListNode
    for curr!=nil {
        next = curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    slow.Next = nil

    // merge the two lists
    first, second := head, prev
    for second!=nil {
        nextFirst, nextSecond := first.Next, second.Next

        first.Next = second
        second.Next = nextFirst

        first = nextFirst
        second = nextSecond
    }
}
