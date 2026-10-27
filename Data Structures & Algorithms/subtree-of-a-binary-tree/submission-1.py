# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        if (p==None) or (q==None):
            return p==q

        if p.val!=q.val:
            return False
        
        isSameTreeLeft = self.isSameTree(p.left, q.left)
        isSameTreeRight = self.isSameTree(p.right, q.right)
        return isSameTreeLeft and isSameTreeRight

    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        if root==None or subRoot==None:
            return root==subRoot

        sameTree = False
        if root.val == subRoot.val:
            sameTree = self.isSameTree(root, subRoot)

        isSubTreeLeft = self.isSubtree(root.left, subRoot)
        isSubTreeRight = self.isSubtree(root.right, subRoot)
        return sameTree or isSubTreeLeft or isSubTreeRight
        