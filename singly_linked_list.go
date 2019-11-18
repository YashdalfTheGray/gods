package gods

import "errors"

// SinglyLinkedList is a list built out nodes that contain
// data and a pointer to the next element in the list.
type SinglyLinkedList struct {
	head   *ListNode
	length uint32
}

// NewSinglyLinkedList created a new instance of a singly
// linked list.
func NewSinglyLinkedList() SinglyLinkedList {
	return SinglyLinkedList{}
}

// Length returns the length of the SinglyLinkedList.
func (ll SinglyLinkedList) Length() uint32 {
	return ll.length
}

// Shift adds an item with the data provided to the front
// of the SinglyLinkedList.
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
// empty list.
func (ll *SinglyLinkedList) Unshift() (interface{}, error) {
	if ll.head == nil {
		return nil, errors.New("The list is empty")
	}

	temp := ll.head
	ll.head = ll.head.Next
	ll.length--
	return temp.Data, nil
}

// Push pushes an item to the end of the SinglyLinkedList.
func (ll *SinglyLinkedList) Push(data interface{}) *SinglyLinkedList {
	if ll.head == nil {
		ll.head = NewListNode(data)
	} else if ll.head.Next == nil {
		ll.head.Next = NewListNode(data)
	} else {
		node := ll.head

		for node.Next != nil {
			node = node.Next
		}

		node.Next = NewListNode(data)
	}

	ll.length++
	return ll
}

// Pop is the opposite of Push. It removed an item from the end
// of a SinglyLinkedList. It also returns an error if the list
// is empty.
func (ll *SinglyLinkedList) Pop() (interface{}, error) {
	if ll.head == nil {
		return nil, errors.New("The list is empty")
	} else if ll.head.Next == nil {
		temp := ll.head
		ll.head = nil
		ll.length--
		return temp.Data, nil
	} else {
		node := ll.head

		for node.Next != nil && node.Next.Next != nil {
			node = node.Next
		}

		temp := node.Next
		node.Next = nil
		ll.length--
		return temp.Data, nil
	}
}

// GetAt returns a peek at the data sitting at a given index.
// This method does not return the item from the list but instead
// a copy of the data contained within the item.
func (ll *SinglyLinkedList) GetAt(index uint32) (interface{}, error) {
	if ll.head == nil {
		return nil, errors.New("The list is empty")
	} else if ll.length <= index {
		return nil, errors.New("Index out of bounds error")
	} else {
		node := ll.head

		for i := uint32(0); i < index; i++ {
			node = node.Next
		}

		return node.Data, nil
	}
}

// Iterate returns an iterator over the contents of the
// SinglyLinkedList.
func (ll SinglyLinkedList) Iterate() Iterator {
	return newSinglyLinkedListIterator(ll)
}

type singlyLinkedListIterator struct {
	ll      SinglyLinkedList
	current *ListNode
}

func newSinglyLinkedListIterator(ll SinglyLinkedList) *singlyLinkedListIterator {
	return &singlyLinkedListIterator{ll: ll}
}

func (i *singlyLinkedListIterator) Next() bool {
	if i.current == nil {
		i.current = i.ll.head
	} else {
		i.current = i.current.Next
	}

	return i.current != nil
}

func (i singlyLinkedListIterator) Get() interface{} {
	return i.current.Data
}

func (i singlyLinkedListIterator) Error() error {
	return nil
}
