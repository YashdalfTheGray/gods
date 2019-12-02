package gods

// TreeNode represents a section of a tree, contains a data
// pointer and its childen
type TreeNode struct {
	Data     interface{}
	Children *SinglyLinkedList
}

// NewTreeNode builds a new tree node and returns it
func NewTreeNode(data interface{}) TreeNode {
	list := NewSinglyLinkedList()
	return TreeNode{data, &list}
}
