package gods

// ListNode represents one node in a linked list
type ListNode struct {
	Data interface{}
	Next *ListNode
	Prev *ListNode
}

// NewListNode will return a ListNode that contains the passed in data
func NewListNode(data interface{}) *ListNode {
	return &ListNode{data, nil, nil}
}
