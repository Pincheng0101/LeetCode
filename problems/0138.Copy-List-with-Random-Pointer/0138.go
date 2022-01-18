package p0138

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	var cachedNode map[*Node]*Node = make(map[*Node]*Node, 0)
	var deepCopy func(*Node) *Node
	deepCopy = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, ok := cachedNode[node]; ok {
			return n
		}
		newNode := &Node{Val: node.Val}
		cachedNode[node] = newNode
		newNode.Next = deepCopy(node.Next)
		newNode.Random = deepCopy(node.Random)
		return newNode
	}

	return deepCopy(head)
}
