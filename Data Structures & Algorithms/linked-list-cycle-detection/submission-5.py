# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        fastPt, slowPt = head, head

        while fastPt!=None :
            slowPt = slowPt.next
            fastPt = fastPt.next
            if fastPt != None:
                fastPt = fastPt.next

            if slowPt!=None and fastPt!=None and slowPt==fastPt:
                return True

        
        return False
        