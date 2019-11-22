package gods

// Stack represents a stack data structure and
// it is based on a singly linked list
type Stack struct {
	list *SinglyLinkedList
}

// NewStack returns a new stack data structure
func NewStack() Stack {
	newList := NewSinglyLinkedList()
	return Stack{list: &newList}
}

// Length returns the current length of the stack
func (s Stack) Length() uint32 {
	return s.list.Length()
}

// Push adds something to the stack
func (s *Stack) Push(data interface{}) *Stack {
	s.list.Push(data)
	return s
}
