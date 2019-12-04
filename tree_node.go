package gods

// TreeNode represents a section of a tree, contains a data
// pointer and its childen
type TreeNode struct {
	Data     interface{}
	Children []*TreeNode
}

// NewTreeNode builds a new tree node and returns it
func NewTreeNode(data interface{}) TreeNode {
	children := make([]*TreeNode, 0)
	return TreeNode{data, children}
}

// IsLeaf returns whether a node is a leaf or not
func (n TreeNode) IsLeaf() bool {
	return len(n.Children) == 0
}
