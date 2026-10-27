/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p==nil || q==nil {
		return p==q
	}

	if p.Val != q.Val {
		return false
	}

	isSameTreeLeft := isSameTree(p.Left, q.Left)
	isSameTreeRight := isSameTree(p.Right, q.Right)
	return isSameTreeLeft && isSameTreeRight
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root==nil || subRoot==nil {
		return root==subRoot
	}

	sameTree := false
	if root.Val==subRoot.Val {
		sameTree = isSameTree(root, subRoot)
	}

	isSubTreeLeft := isSubtree(root.Left, subRoot)
	isSubTreeRight := isSubtree(root.Right, subRoot)
	return sameTree || isSubTreeLeft || isSubTreeRight 
}
