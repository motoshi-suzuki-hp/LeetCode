/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func bstToGst(root *TreeNode) *TreeNode {
    btg(root, 0)
    return root
}

func btg(tn *TreeNode, sum int) int {
    if tn == nil {
        return sum
    }

    sum = btg(tn.Right, sum)
    sum += tn.Val
    tn.Val = sum

    return btg(tn.Left, sum)
}
