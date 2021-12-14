package dsll

import dsnode "go_data_structures/data_structures/node"

type CircLinkedList struct {
	Head *dsnode.Node
	Tail *dsnode.Node
}

func NewCircLinkedList() *CircLinkedList {
	var newCLL *CircLinkedList = new(CircLinkedList)

	newCLL.Head = nil
	newCLL.Tail = nil

	return newCLL
}

func (CLL *CircLinkedList) Clone() {}

func Insert(index int, val int) {}
