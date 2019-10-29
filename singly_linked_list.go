package gods

import "errors"

// SinglyLinkedList is a list built out nodes that contain
//  data and a pointer to the next element in the list
type SinglyLinkedList struct {
	head   *ListNode
	length uint32
}

// NewSinglyLinkedList created a new instance of a singly
// linked list.
func NewSinglyLinkedList() SinglyLinkedList {
	return SinglyLinkedList{}
}

// Length returns the length of the SinglyLinkedList
func (ll SinglyLinkedList) Length() uint32 {
	return ll.length
}

// Shift adds an item with the data provided to the front
// of the SinglyLinkedList
func (ll *SinglyLinkedList) Shift(data interface{}) *SinglyLinkedList {
	if ll.head == nil {
		ll.head = NewListNode(data)
		ll.head.Next = nil
	} else {
		temp := ll.head
		ll.head = NewListNode(data)
		ll.head.Next = temp
	}

	ll.length++
	return ll
}

// Unshift is the opposite of Shift. It removed one item from
// the start of the list. It also returns an error for an
// empty list
func (ll *SinglyLinkedList) Unshift() (interface{}, error) {
	if ll.head == nil {
		return nil, errors.New("The list is empty")
	}

	temp := ll.head
	ll.head = ll.head.Next
	ll.length--
	return temp.Data, nil
}

// Iterate returns an iterator over the contents of the
// SinglyLinkedList
func (ll SinglyLinkedList) Iterate() Iterator {
	return newLinkedListIterator(ll)
}

type linkedListIterator struct {
	ll      SinglyLinkedList
	current *ListNode
}

func newLinkedListIterator(ll SinglyLinkedList) *linkedListIterator {
	return &linkedListIterator{ll: ll}
}

func (i *linkedListIterator) Next() bool {
	if i.current == nil {
		i.current = i.ll.head
	} else {
		i.current = i.current.Next
	}

	return i.current != nil
}

func (i linkedListIterator) Get() interface{} {
	return i.current.Data
}

func (i linkedListIterator) Error() error {
	return nil
}
