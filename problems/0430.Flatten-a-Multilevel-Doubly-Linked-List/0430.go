package p0430

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	dfs(root)
	return root
}

// 回傳最後一個 Node
func dfs(root *Node) (last *Node) {
	cur := root

	// 如果有子節點，優先處理
	for cur != nil {
		next := cur.Next
		if cur.Child != nil {
			childLastNode := dfs(cur.Child)

			// 將 cur Node 與 cur.Child Node 相連
			connect(cur, cur.Child)

			// 如果 next 不為空，就將 childLastNode 與 next Node 相連
			if next != nil {
				connect(childLastNode, next)
			}

			cur.Child = nil
			last = childLastNode
		} else {
			last = cur
		}
		cur = next
	}
	return
}

// 連接兩個節點
func connect(a *Node, b *Node) {
	a.Next, b.Prev = b, a
}
