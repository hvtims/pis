package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil && l2.Head == nil {
		return
	} else if l1.Head == nil {
		*l1 = *l2
		return
	} else if l2.Head == nil {
		return
	}
	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}
