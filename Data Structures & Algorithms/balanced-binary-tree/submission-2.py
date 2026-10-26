# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def __init__(self):
        self.isTreeBalanced = True

    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        def treeTraversal(node: Optional[TreeNode]) -> int:
            if node==None :
                return 0

            leftHeight = treeTraversal(node.left)
            rightHeight = treeTraversal(node.right)
            if abs(leftHeight-rightHeight)>1:
                self.isTreeBalanced = False

            return 1 + max(leftHeight, rightHeight)

        treeTraversal(root)
        return self.isTreeBalanced