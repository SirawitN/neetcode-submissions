/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    nNodes := 0
    curNode := head
    for curNode!=nil {
        nNodes++
        curNode = curNode.Next
    }

    var prevNode *ListNode
    curNode = head
    i := 0
    for {
        if i==nNodes-n {
            if prevNode==nil {
                head = head.Next
                curNode.Next = nil
            } else {
                prevNode.Next = curNode.Next
                curNode.Next = nil
            }
            break
        }
        prevNode = curNode
        curNode = curNode.Next
        i++
    }

    return head
}
