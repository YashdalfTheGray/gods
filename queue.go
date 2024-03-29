package gods

import "errors"

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

// Enqueue adds stuff to the end of the queue
func (q *Queue) Enqueue(data interface{}) *Queue {
	q.list.Push(data)
	return q
}

// Dequeue removes stuff from the start of the queue
func (q *Queue) Dequeue() (interface{}, error) {
	val, err := q.list.Unshift()
	if err != nil {
		return nil, errors.New("The queue is empty")
	}
	return val, nil
}

// Peek returns a copy of the element at the start
// of the queue
func (q Queue) Peek() (interface{}, error) {
	val, err := q.list.GetAt(0)
	if err != nil {
		return nil, errors.New("The queue is empty")
	}
	return val, nil
}

// IsEmpty returns whether the queue is empty
func (q Queue) IsEmpty() bool {
	return q.list.Length() == 0
}
