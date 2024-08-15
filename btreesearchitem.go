package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}

	if leftResult := BTreeSearchItem(root.Left, elem); leftResult != nil {
		return leftResult
	}

	return BTreeSearchItem(root.Right, elem)
}
