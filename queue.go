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
