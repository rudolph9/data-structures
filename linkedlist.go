package linkedlist


type Node struct {
	value string
	next *Node
}

type LinkedList struct {
	head *Node
}

// returns the inserted node
func (n *Node) InsertNext(val string) (*Node){
	next := &Node{ value: val }
	n.next = next
	return next
}

func (n *Node) InsertPrev(val string) (*Node){
	return &Node{ value: val, next: n }
}

func (l LinkedList) Print() (string){
	var rtn string
	cur := l.head
	for {
		if cur == nil {
			break
		}
		rtn += cur.value + "\n"
		cur = cur.next
	}

	return rtn
}







