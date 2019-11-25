package gods

import "errors"

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

// Pop removed the top element from the stack
func (s *Stack) Pop() (interface{}, error) {
	val, err := s.list.Pop()
	if err != nil {
		return nil, errors.New("The stack is empty")
	}
	return val, nil
}

// Peek returns a copy of the element at the top
// of the stack but doesn't pop the element
func (s *Stack) Peek() (interface{}, error) {
	val, err := s.list.GetAt(s.list.Length() - 1)
	if err != nil {
		return nil, errors.New("The stack is empty")
	}
	return val, nil
}

// IsEmpty returns whether the stack is empty
func (s Stack) IsEmpty() bool {
	return s.list.Length() == 0
}
