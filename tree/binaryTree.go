package tree

func (bt *BinaryTree) Search(number int) bool {
	isFound := false

	currentNode := bt.Root
	for {
		if number == currentNode.Value {
			isFound = true
			break
		}

		if number < currentNode.Value {
			if isNodeNil(currentNode.LeftNode) {
				break
			}
			currentNode = currentNode.LeftNode
		}

		if number > currentNode.Value {
			if isNodeNil(currentNode.RightNode) {
				break
			}
			currentNode = currentNode.RightNode
		}
	}

	return isFound
}

func isNodeNil(node *Node) bool {
	return node == nil
}

func NewBinaryTree(items []int) *BinaryTree {
	binaryTree := &BinaryTree{
		Root: &Node{},
	}

	treeItems := OrderSliceAndRemoveDuplicates(items)

	pivotPosition := GetPivotPosition(treeItems)

	binaryTree.Root.LeftNode = getTreeNode(treeItems, pivotPosition, getLeftNodeItems)
	binaryTree.Root.RightNode = getTreeNode(treeItems, pivotPosition, getRightNodeItems)

	binaryTree.Root.Value = treeItems[pivotPosition]

	return binaryTree
}

func GetPivotPosition(slice []int) int {
	pivotPosition := len(slice) / 2
	return int(pivotPosition)
}

func getTreeNode(parentItems []int, parentPivotPosition int, getNodeItemsFunc NodeItemsFunc) *Node {
	node := &Node{}

	nodeItems := getNodeItemsFunc(parentItems, parentPivotPosition)
	nodePivotPosition := GetPivotPosition(nodeItems)

	node.Value = nodeItems[nodePivotPosition]

	if hasChildNode(nodeItems, LEFT_NODE) {
		node.LeftNode = getTreeNode(nodeItems, nodePivotPosition, getLeftNodeItems)
	}

	if hasChildNode(nodeItems, RIGHT_NODE) {
		node.RightNode = getTreeNode(nodeItems, nodePivotPosition, getRightNodeItems)
	}

	return node
}

func getLeftNodeItems(treeItems []int, pivotPosition int) []int {
	return treeItems[:pivotPosition]
}

func getRightNodeItems(treeItems []int, pivotPosition int) []int {
	return treeItems[pivotPosition+1:]
}

func hasChildNode(nodeItems []int, nodeType NodeType) bool {

	if nodeType == LEFT_NODE {
		return len(nodeItems) > 1
	}

	if nodeType == RIGHT_NODE {
		return len(nodeItems) > 2
	}

	return false
}
