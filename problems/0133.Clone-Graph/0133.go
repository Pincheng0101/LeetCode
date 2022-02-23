package p0133

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	clonedMap := map[int]*Node{}
	return dfs(node, clonedMap)
}

func dfs(node *Node, clonedMap map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	if _, ok := clonedMap[node.Val]; ok {
		return clonedMap[node.Val]
	}

	newNode := &Node{Val: node.Val}
	clonedMap[node.Val] = newNode
	for _, neighborNode := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, dfs(neighborNode, clonedMap))
	}

	return newNode
}
