# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def __init__(self):
        self.lca: TreeNode = None

    def dfs(self, node: TreeNode, p: TreeNode, q: TreeNode) -> bool:
        if node==None:
            return False
        
        selfLca = (node.val==p.val) or (node.val==q.val)
        leftLca = self.dfs(node.left, p, q)
        rightLca = self.dfs(node.right, p, q)

        cond1 = selfLca and (leftLca or rightLca)
        cond2 = leftLca and rightLca

        if selfLca or leftLca or rightLca or cond1 or cond2:
            self.lca = node if ((cond1 or cond2) and not self.lca) else self.lca
            return True

        return False

    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        self.dfs(root, p, q)
        return self.lca
        