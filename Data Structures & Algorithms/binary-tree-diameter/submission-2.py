# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def __init__(self):
        self.diameter = 0

    def findDiameter(self, node: Optional[TreeNode]) -> int:
        if node==None:
            return 0

        leftHeight = self.findDiameter(node.left)
        rightHeight = self.findDiameter(node.right)
        self.diameter = max(self.diameter, leftHeight+rightHeight)

        return 1 + max(leftHeight, rightHeight)
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        self.findDiameter(root)

        return self.diameter