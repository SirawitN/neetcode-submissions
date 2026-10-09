# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
	def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
		if list1==None and list2==None:
			return None

		if list1!=None and (list2==None or list1.val <= list2.val):
			sortedList = list1
			currNode = list1
			list1 = list1.next
		else:
			sortedList = list2
			currNode = list2
			list2 = list2.next

		while list1!=None or list2!=None:
			if list1==None:
				currNode.next = list2
				list2 = list2.next
			elif list2==None:
				currNode.next = list1
				list1 = list1.next
			else:
				if list1.val <= list2.val:
					currNode.next = list1
					list1 = list1.next
				else:
					currNode.next = list2
					list2 = list2.next
			currNode = currNode.next

		return sortedList
        