package gods

import "errors"

// Direction represents a direction for the iterator
type Direction int

// Constants for each of the directions that we can iterate
const (
	Forward Direction = iota
	Reverse
)

// DoublyLinkedList is a type of list with nodes
// that hold references to the node before and the
// node after.
type DoublyLinkedList struct {
	head     *ListNode
	tail     *ListNode
	length   uint32
	circular bool
}

// NewDoublyLinkedList created a new list and returns it.
// Passing in true to the constructor will make it a circular
// list, i.e., the last node will point back to the first node
func NewDoublyLinkedList(circular bool) DoublyLinkedList {
	return DoublyLinkedList{circular: circular}
}

// Length returns the total number of nodes in the list
func (ll DoublyLinkedList) Length() uint32 {
	return ll.length
}

// Shift adds an item that contains the given data to the start
// of the DoublyLinkedList
func (ll *DoublyLinkedList) Shift(data interface{}) *DoublyLinkedList {
	if ll.head == nil {
		ll.head = NewListNode(data)
		ll.tail = ll.head
	} else {
		temp := ll.head
		ll.head = NewListNode(data)

		ll.head.Next = temp
		temp.Prev = ll.head

		if ll.circular {
			ll.head.Prev = ll.tail
			if ll.tail != nil {
				ll.tail.Next = ll.head
			}
		}
	}

	ll.length++
	return ll
}

// Unshift is the opposite of Shift. It removed one item from
// the start of the list. It also returns an error for an
// empty list.
func (ll *DoublyLinkedList) Unshift() (interface{}, error) {
	if ll.head == nil {
		return nil, errors.New("The list is empty")
	} else if ll.length == 1 {
		temp := ll.head
		ll.head = nil
		ll.tail = nil
		return temp.Data, nil
	}

	temp := ll.head
	ll.head = ll.head.Next

	if ll.circular {
		ll.head.Prev = ll.tail
		if ll.tail != nil {
			ll.tail.Next = ll.head
		}
	} else {
		ll.head.Prev = nil
	}

	ll.length--
	return temp.Data, nil
}

// Push pushes an item to the end of the DoublyLinkedList.
func (ll *DoublyLinkedList) Push(data interface{}) *DoublyLinkedList {
	if ll.head == nil {
		ll.head = NewListNode(data)
		ll.tail = ll.head
	} else {
		temp := ll.tail
		ll.tail = NewListNode(data)

		ll.tail.Prev = temp
		temp.Next = ll.tail

		if ll.circular {
			ll.tail.Next = ll.head
		}
	}

	ll.length++
	return ll
}

// Pop is the opposite of Push. It removed an item from the end
// of a SinglyLinkedList. It also returns an error if the list
// is empty.
func (ll *DoublyLinkedList) Pop() (interface{}, error) {
	if ll.head == nil {
		return nil, errors.New("The list is empty")
	} else if ll.head.Next == nil {
		temp := ll.head
		ll.head = nil
		ll.length--
		return temp.Data, nil
	} else {
		temp := ll.tail
		ll.tail = ll.tail.Prev

		if ll.circular {
			ll.head.Prev = ll.tail
			ll.tail.Next = ll.head
		} else {
			ll.tail.Next = nil
		}

		ll.length--
		return temp.Data, nil
	}
}

// Iterate returns an iterator over the contents of the
// DoublyLinkedList. It can iterate in either direction.
func (ll DoublyLinkedList) Iterate(d Direction) Iterator {
	return newDoublyLinkedListIterator(ll, d)
}

type doublyLinkedListIterator struct {
	ll        DoublyLinkedList
	current   *ListNode
	goal      *ListNode
	direction Direction
}

func newDoublyLinkedListIterator(ll DoublyLinkedList, d Direction) *doublyLinkedListIterator {
	return &doublyLinkedListIterator{ll: ll, direction: d}
}

func (i *doublyLinkedListIterator) Next() bool {
	if i.current == nil {
		if i.direction == Forward {
			i.current = i.ll.head
			i.goal = i.ll.tail
		} else {
			i.current = i.ll.tail
			i.goal = i.ll.head
		}
	} else {
		if i.direction == Forward {
			i.current = i.current.Next
		} else {
			i.current = i.current.Prev
		}
	}

	if i.ll.circular {
		return i.current != i.goal
	}
	return i.current != nil

}

func (i doublyLinkedListIterator) Get() interface{} {
	return i.current.Data
}

func (i doublyLinkedListIterator) Error() error {
	return nil
}
