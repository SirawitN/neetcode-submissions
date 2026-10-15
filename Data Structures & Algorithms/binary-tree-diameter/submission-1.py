# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
	def __init__(self):
		self.maxDiameter = 0

	def treeHeight(self, node: TreeNode) -> int:
		if node==None:
			return 0

		leftTreeHeight = self.treeHeight(node.left)
		rightTreeHeight = self.treeHeight(node.right)
		self.maxDiameter = max(self.maxDiameter, leftTreeHeight + rightTreeHeight)

		return 1 + max(leftTreeHeight, rightTreeHeight)

	def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
		self.treeHeight(root)

		return self.maxDiameter

        