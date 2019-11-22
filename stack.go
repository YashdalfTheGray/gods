package gods

// Stack represents a stack data structure and
// it is based on a singly linked list
type Stack struct {
	list *SinglyLinkedList
}

// NewStack returns a new stack data structure
func NewStack() Stack {
	return Stack{}
}
