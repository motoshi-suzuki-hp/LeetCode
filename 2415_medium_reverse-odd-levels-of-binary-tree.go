/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func reverseOddLevels(root *TreeNode) *TreeNode {
    if root == nil {
		return nil
	}
    dfs(root.Left, root.Right, 1)
    return root
}

func dfs(l, r *TreeNode, level int) {
    if l == nil || r == nil {
        return
    }
    if level%2 == 1 {
        l.Val, r.Val = r.Val, l.Val
    }
    dfs(l.Left, r.Right, level+1)
    dfs(r.Left, l.Right, level+1)
}
