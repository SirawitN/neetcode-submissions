/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
        return
    }

    // Step 1: Find the middle of the list
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // Step 2: Reverse the second half of the list
    var prev *ListNode
    curr := slow.Next
    slow.Next = nil // Sever the first half from the second half
    
    for curr != nil {
        nextTemp := curr.Next
        curr.Next = prev
        prev = curr
        curr = nextTemp
    }
    // 'prev' is now the head of the reversed second half

    // Step 3: Merge the two halves
	reordered := head
    first, second := head, prev
    for second != nil {
        // Save next pointers
        nextFirst := first.Next
        nextSecond := second.Next

        // Reassign connections
        first.Next = second
        second.Next = nextFirst

        // Move pointers forward
        first = nextFirst
        second = nextSecond
    }

	head = reordered
}
