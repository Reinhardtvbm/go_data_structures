package main

import (
	"fmt"
	// ds short for data structures
	ds "go_data_structures/data_structures/node"
)

func main() {
	var testNode *(ds.Node) = ds.NewNode(1, nil)
	fmt.Println(testNode.Val)
}
