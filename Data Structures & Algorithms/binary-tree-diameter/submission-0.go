/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0

	var diameter func(node *TreeNode) int
	diameter = func(node *TreeNode) int {
		if node==nil {
			return 0
		}

		leftHeight := diameter(node.Left)
		rightHeight := diameter(node.Right)
		maxDiameter = max(maxDiameter, leftHeight+rightHeight)

		return 1 + max(leftHeight, rightHeight)
	} 

	diameter(root)
	return maxDiameter
}
