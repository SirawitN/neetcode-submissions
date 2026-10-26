/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    isBalance := true

	var treeTraversal func(node *TreeNode) float64
	treeTraversal = func(node *TreeNode) float64 {
		if node==nil {
			return 0.0
		}

		leftHeight := treeTraversal(node.Left)
		rightHeight := treeTraversal(node.Right)
		if math.Abs(leftHeight - rightHeight) > 1 {
			isBalance = false
		}

		return 1 + max(leftHeight, rightHeight)
	}

	treeTraversal(root)
	return isBalance
}
