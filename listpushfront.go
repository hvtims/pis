package piscine

func ListPushFront(l *List, data interface{}) {
	Newnode := &NodeL{Data: data}
	ns := l.Head
	if l.Tail == nil {
		l.Head = Newnode
		l.Tail = Newnode
	} else {
		l.Head = Newnode
		l.Head.Next = ns
	}
}
