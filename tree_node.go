package gods

// TreeNode represents a section of a tree, contains a data
// pointer and its childen
type TreeNode struct {
	data     interface{}
	Children []TreeNode
}

// NewTreeNode builds a new tree node and returns it
func NewTreeNode(data interface{}) TreeNode {
	return TreeNode{data, []TreeNode{}}
}
