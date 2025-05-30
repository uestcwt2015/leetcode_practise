/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
var cacheNodes map[*Node]*Node
func deepCopy(node *Node) *Node {
    if node == nil {
        return nil
    }

    if val, has := cacheNodes[node]; has {
        return val
    }

    newNode := &Node{Val: node.Val}
    cacheNodes[node] = newNode
    newNode.Next = deepCopy(node.Next)
    newNode.Random = deepCopy(node.Random)

    return newNode
}

func copyRandomList(head *Node) *Node {
    cacheNodes = map[*Node]*Node{}
    return deepCopy(head)
}