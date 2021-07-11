func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 查找节点所在位置 找到节点返回非 nil
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	// 遍历左右子树 查找节点是否存在
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 左右节点都存在
	if left != nil && right != nil {
		return root
	}

	// 返回存在的节点
	if left == nil {
		return right
	}
	return left
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/solution/er-cha-shu-de-zui-jin-gong-gong-zu-xian-by-leetc-2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
