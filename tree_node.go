package gods

import "errors"

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

// AddChild adds a child to the tree node
func (n *TreeNode) AddChild(child interface{}) *TreeNode {
	newTreeNode := NewTreeNode(child)
	n.Children = append(n.Children, &newTreeNode)

	return n
}

// RemoveChild removes a child from the tree node. Any children
// of the removed node become children of this tree node
func (n *TreeNode) RemoveChild(childData interface{}) (interface{}, error) {
	if len(n.Children) == 0 {
		return nil, errors.New("This node does not have any children")
	}

	index := 0
	for i, child := range n.Children {
		if child.Data == childData {
			index = i
		}
	}

	childrenBefore := n.Children[0 : index-1]
	removedChild := n.Children[index]
	childrenAfter := n.Children[index+1 : len(n.Children)-1]

	if !removedChild.IsLeaf() {
		childrenBefore = append(childrenBefore, removedChild.Children...)
	}
	n.Children = append(childrenBefore, childrenAfter...)

	return removedChild.Data, nil
}
