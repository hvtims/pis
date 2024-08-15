package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	current := l.Head
	c := 0
	for current != nil {
		c++
		current = current.Next
	}
	return c
}
