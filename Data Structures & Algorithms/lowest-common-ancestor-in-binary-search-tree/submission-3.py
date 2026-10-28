# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def __init__(self):
        self.lca: TreeNode = None
    def dfs(self, node: TreeNode, p: TreeNode, q: TreeNode):
        if node==None:
            return False

        selfAncestor = (node.val==p.val) or (node.val==q.val)
        leftLCA = self.dfs(node.left, p, q)
        rightLCA = self.dfs(node.right, p, q)
        
        cond1 = selfAncestor and (leftLCA or rightLCA)
        cond2 = leftLCA and rightLCA

        if cond1 or cond2 or selfAncestor or leftLCA or rightLCA:
            self.lca = node if (cond1 or cond2) and self.lca==None else self.lca
            return True
        return False
    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        self.dfs(root, p, q)
        return self.lca