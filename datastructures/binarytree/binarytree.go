package binarytree

import "container/list"

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

// [4,2,9,3,5,null,7] =>
//        4
//      2   9
//     3 5   7
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

func (t *TreeNode) GetNodeValuesWithDFS() []int {
	values := []int{}
	var dfs func(*TreeNode)
	dfs = func(td *TreeNode) {
		if td == nil {
			return
		}
		values = append(values, td.Val)
		dfs(td.Left)
		dfs(td.Right)
	}
	dfs(t)

	return values
}

func (t *TreeNode) GetNodeValuesWithBFS() []int {
	values := []int{}
	queue := list.New()
	queue.PushBack(t)
	var bfs func()
	bfs = func() {
		if queue.Len() == 0 {
			return
		}

		for i := 0; i < queue.Len(); i++ {
			qnode := queue.Front()
			node := qnode.Value.(*TreeNode)
			values = append(values, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			queue.Remove(qnode)
		}
		bfs()
	}
	bfs()

	return values
}
