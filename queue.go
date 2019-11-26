package gods

// Queue is a simple queue data structure backed by
// a singly linked list
type Queue struct {
	list SinglyLinkedList
}

// NewQueue returns a new instance of the Queue data
// structure
func NewQueue() *Queue {
	return &Queue{}
}

// Size returns the size of the queue itself
func (q Queue) Size() uint32 {
	return q.list.Length()
}
