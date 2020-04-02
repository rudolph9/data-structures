package linkedlist

import "testing"

func TestInsertNext(t *testing.T) {
	myNode := Node{
		value: "myNode",
	}

	myNode.InsertNext("myNode1")

	if myNode.next.value != "myNode1" {
		t.Errorf("next is not the inserted node")
	}
}

func TestInsertPrev(t *testing.T) {
	myNode1 := Node{
		value: "myNode1",
	}

	myNode := myNode1.InsertPrev("myNode")

	if myNode.next.value != "myNode1" {
		t.Errorf("next is not the inserted node")
	}
}
