package linkedlist

type value string

type Node struct {
	value
	next *Node
}

type LinkedList struct {
	head *Node
}

// returns the inserted node
func (n *Node) InsertNext(val value) (*Node){
	next := &Node{ value: val }
	n.next = next
	return next
}

func (n *Node) InsertPrev(val value) (*Node){
	return &Node{ value: val, next: n }
}





