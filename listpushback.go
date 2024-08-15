package piscine

func ListPushBack(l *List, data interface{}) {
	newNode := &NodeL{Data: data}
	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
