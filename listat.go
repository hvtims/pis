package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	current := l
	for pos != 0 {
		pos--
		if current.Next == nil {
			return nil
		}
		current = current.Next
	}
	return current
}
