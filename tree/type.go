package tree

const (
	LEFT_NODE  NodeType = "LEFT"
	RIGHT_NODE NodeType = "RIGHT"
)

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Value     int
	LeftNode  *Node
	RightNode *Node
}

type NodeItemsFunc func([]int, int) []int

type NodeType string
