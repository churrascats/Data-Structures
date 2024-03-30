package tree

type BinaryTree struct {
	Root Node
}

type Node struct {
	Value     int
	LeftNode  *Node
	RightNode *Node
}

type NodeBranchFunction func([]int, int) []int

func NewBinaryTreeFromSlice(treeItems []int) {
	newTree := &BinaryTree{}

	pivotPosition := GetPivotPosition(treeItems)

	newTree.Root.Value = treeItems[pivotPosition]
	setBinaryTreeNodesFromSlice(&newTree.Root, getLeftNodeSlice, treeItems)
	setBinaryTreeNodesFromSlice(&newTree.Root, getRightNodeSlice, treeItems)
}

func GetPivotPosition(slice []int) int {
	pivotPosition := len(slice) / 2
	return int(pivotPosition)
}

func setBinaryTreeNodesFromSlice(node *Node, nodeBranchFunction NodeBranchFunction, treeItems []int) {
	nodePivotPosition := GetPivotPosition(treeItems)
	nodeItems := nodeBranchFunction(treeItems, nodePivotPosition)

	if len(nodeItems) > 0 {
		node.LeftNode.Value = nodeItems[nodePivotPosition]
	}

	if nodePivotPosition > 0 {
		setBinaryTreeNodesFromSlice(node, nodeBranchFunction, nodeItems)
	}

}

func getLeftNodeSlice(treeItems []int, pivotPosition int) []int {
	return treeItems[:pivotPosition]
}

func getRightNodeSlice(treeItems []int, pivotPosition int) []int {
	return treeItems[pivotPosition+1:]
}
