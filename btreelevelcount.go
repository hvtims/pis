package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	right := BTreeLevelCount(root.Right)
	Left := BTreeLevelCount(root.Left)
	if Left > right {
		return 1 + Left
	} else {
		return 1 + right
	}
}
