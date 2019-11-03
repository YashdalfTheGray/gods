package gods

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
