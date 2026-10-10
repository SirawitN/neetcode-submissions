# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
	def reorderList(self, head: Optional[ListNode]) -> None:
		lists = list()
		while head!=None:
			lists.append(head)
			head = head.next

		head, currNode = lists[0], lists[0]
		left, right = min(1, len(lists)-1), len(lists)-1

		while left<=right:
			currNode.next = lists[right]
			currNode.next.next = lists[left]

			currNode = currNode.next.next
			left += 1
			right -= 1

		currNode.next = None