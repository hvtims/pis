package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return BTreeHelper(root, nil, nil)
}

func BTreeHelper(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if max != nil && root.Data > max.Data || min != nil && root.Data < min.Data {
		return false
	}
	return BTreeHelper(root.Left, min, root) && BTreeHelper(root.Right, root, max)
}
