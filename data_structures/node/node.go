package ds

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int, next *Node) *Node {
	var newNode *Node = new(Node)

	newNode.Val = val
	newNode.Next = next

	return newNode
}
