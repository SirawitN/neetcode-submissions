/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    var LCA *TreeNode
	var dfs func(node *TreeNode, p *TreeNode, q *TreeNode) bool
	dfs = func(node *TreeNode, p *TreeNode, q *TreeNode) bool {
		if node==nil {
			return false
		}
		selfAncestor := (node.Val==p.Val) || (node.Val==q.Val)

		leftLCA := dfs(node.Left, p, q)
		rightLCA := dfs(node.Right, p, q)

		// fmt.Println(node.Val, p.Val, q.Val, selfAncestor, leftLCA, rightLCA)

		cond1 := selfAncestor && (leftLCA || rightLCA)
		cond2 := leftLCA && rightLCA
		
		if (cond1 || cond2) && LCA==nil {
			LCA = node
			return true
		} else if selfAncestor || leftLCA || rightLCA {
			return true
		}
		return false
	}

	dfs(root, p, q)
	return LCA
}
