package dsll

import (
	"fmt"
	dsnode "go_data_structures/data_structures/node"
	"log"
)

type LinkedList struct {
	Head *dsnode.Node
	Tail *dsnode.Node
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

	var traverse *dsnode.Node = other.Head

	for traverse != nil {
		clone.Add(traverse.Val)
	}

	return clone
}

// add a new item to end of the list
func (LL *LinkedList) Add(newVal int) {
	// create new node
	var newNode *dsnode.Node = dsnode.NewNode(newVal, nil)

	// if empty list make head new node
	if LL.Head == nil {
		LL.Head = newNode
		LL.Tail = LL.Head
	}

	LL.Tail.Next = newNode
	LL.Tail = newNode // have last Node point to new node
}

func (LL *LinkedList) Insert(index int, val int) {
	size := LL.Size()

	if index <= size && index >= 0 {
		if index == size {
			LL.Add(val) // if adding to last index, just add to the end
		} else {
			var traverse *dsnode.Node = LL.Head

			for i := 0; i < index; i++ {
				traverse = traverse.Next
			}

			var newNode *dsnode.Node = dsnode.NewNode(val, traverse.Next) // the node to be inserted
			traverse.Next = newNode
		}
	} else {
		log.Println("Error: index out of bounds")
	}
}

func (LL *LinkedList) Remove(index int) {
	if LL.Head != nil {
		if index < LL.Size() && index >= 0 {
			if LL.Size() == 1 {
				LL.Head = nil
				LL.Tail = nil
			} else if index == 0 {
				LL.Head = LL.Head.Next
			} else {
				var curr *dsnode.Node = LL.Head.Next
				var prev *dsnode.Node = LL.Head

				for i := 0; i < index-1; i++ {
					prev = curr
					curr = curr.Next
				}

				prev.Next = curr.Next
			}
		} else {
			fmt.Print("Error: index out of bounds!\n")
			return
		}
	} else {
		fmt.Print("Error: the linkedList is empty!\n")
	}

}

func (LL *LinkedList) IsEmpty() bool {
	return LL.Head == nil
}

func (LL *LinkedList) Clear() {
	LL.Head = nil
	LL.Tail = nil
}

func (LL *LinkedList) GetLeader() *dsnode.Node {
	return LL.Head
}

func (LL *LinkedList) Size() int {
	if LL.Head == nil {
		return 0
	}

	var traverse *dsnode.Node = LL.Head
	var count int = 0

	for traverse != nil {
		count++
		traverse = traverse.Next
	}

	return count
}

func (LL *LinkedList) PrintList() {
	// start traversing the list at the head
	var traverse *dsnode.Node = LL.Head

	if traverse == nil {
		fmt.Print("LinkedList: []\n")
		return
	}

	fmt.Print("LinkedList: [")

	for traverse.Next != nil {
		fmt.Print(traverse.Val, ",")
		traverse = traverse.Next
	}

	fmt.Print(traverse.Val, "]\n")
}
