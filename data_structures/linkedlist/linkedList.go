package ds

import ds "go_data_structures/data_structures/node"

type LinkedList struct {
	Head *ds.Node
	Tail *ds.Node
}

func NewLinkedList() *LinkedList {
	var LL *LinkedList = new(LinkedList)

	LL.Head = nil
	LL.Tail = nil

	return LL
}

// return deep copy
func Clone(other *LinkedList) *LinkedList {
	var clone *LinkedList = NewLinkedList()

	var traverse *ds.Node = other.Head

	for traverse != nil {
		clone.Add(traverse.Val)
	}

	return clone
}

// add a new item to end of the list
func (LL *LinkedList) Add(newVal int) {
	// create new node
	var newNode *ds.Node = ds.NewNode(newVal, nil)

	// if empty list make head new node
	if LL.Head == nil {
		LL.Head = newNode
		LL.Tail = LL.Head
	}

	LL.Tail.Next = newNode
	LL.Tail = newNode // have last Node point to new node
}

func (LL *LinkedList) Insert(index int, val int) {

}

func (LL *LinkedList) Remove(index int) {

}

func (LL *LinkedList) IsEmpty() bool {
	if LL.Head == nil {
		return true
	} else {
		return false
	}
}

func (LL *LinkedList) Clear() {
	LL.Head = nil
	LL.Tail = nil
}

func (LL *LinkedList) GetLeader() *ds.Node {
	return LL.Head
}
