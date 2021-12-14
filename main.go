package main

import (
	"fmt"
	dsll "go_data_structures/data_structures/linked_list"
)

func main() {
	var LL1 *dsll.LinkedList = dsll.NewLinkedList()
	LL1.Add(1)
	LL1.Add(2)
	LL1.Add(3)
	LL1.Add(4)
	LL1.Add(5)
	LL1.Add(6)
	LL1.Add(7)
	LL1.Add(8)
	LL1.PrintList()

	var LL2 *dsll.LinkedList = LL1.Clone()
	LL2.PrintList()

	fmt.Println("LL1 address: ", LL1)
	fmt.Println("LL2 address: ", LL2)

	LL2.Insert(3, 3)
	LL2.Insert(9, 9)
	LL2.Insert(12, 9)
	LL2.PrintList()
}
