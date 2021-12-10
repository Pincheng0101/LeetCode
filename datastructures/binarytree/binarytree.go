package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func NewTreeNode() *TreeNode {
	return &TreeNode{}
}

func NewBinaryTree(val []interface{}) *BinaryTree {
	tree := BinaryTree{}
	tree.Root = AddNode(val)
	return &tree
}

func AddNode(val []interface{}) *TreeNode {
	if len(val) == 0 || val[0] == nil {
		return nil
	}
	node := NewTreeNode()
	node.Val = val[0].(int)
	leftNodeVal := make([]interface{}, 0)
	rightNodeVal := make([]interface{}, 0)
	layerNodeCount := 1
	// left: 1,3,4,7,8,9,10...
	// right: 2,5,6,11,12,13,14...
	length := len(val)
	for i := 1; i < length; {
		layerOffset := i
		for ; i < layerOffset+layerNodeCount && i < length; i++ {
			leftNodeVal = append(leftNodeVal, val[i])
		}
		for ; i < layerOffset+layerNodeCount*2 && i < length; i++ {
			rightNodeVal = append(rightNodeVal, val[i])
		}
		layerNodeCount *= 2
	}
	node.Left = AddNode(leftNodeVal)
	node.Right = AddNode(rightNodeVal)
	return node
}
