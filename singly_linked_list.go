package gods

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
