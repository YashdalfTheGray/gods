package gods_test

import (
	"testing"

	"github.com/YashdalfTheGray/gods"
)

func TestNewQueue(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "Creates a new queue",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if q := gods.NewQueue(); q.Size() != 0 {
				t.Errorf("Expected length to be 0 but got %d", q.Size())
			}
		})
	}
}

func TestQueueSize(t *testing.T) {
	testCases := []struct {
		desc    string
		content []int
		length  uint32
	}{
		{
			desc:    "returns 0 for empty queue",
			content: []int{},
			length:  0,
		},
		{
			desc:    "returns 1 for a queue with 1 item in it",
			content: []int{1},
			length:  1,
		},
		{
			desc:    "returns length of queue",
			content: []int{1, 2, 3, 4},
			length:  4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			q := gods.NewQueue()

			for _, v := range tC.content {
				q.Enqueue(v)
			}

			if qLength := q.Size(); qLength != tC.length {
				t.Errorf("Expected length to be %d but got %d", tC.length, qLength)
			}
		})
	}
}

func TestQueueEnqueue(t *testing.T) {
	testCases := []struct {
		desc    string
		in1     string
		in2     string
		in3     string
		addAll  bool
		prepend []string
		length  uint32
	}{
		{
			desc:    "can push something on an empty queue",
			in1:     "things",
			in2:     "",
			in3:     "",
			addAll:  false,
			prepend: []string{},
			length:  1,
		},
		{
			desc:    "can push something on a non-empty queue",
			in1:     "bar",
			in2:     "",
			in3:     "",
			addAll:  false,
			prepend: []string{"foo"},
			length:  2,
		},
		{
			desc:    "supports fluent API",
			in1:     "things",
			in2:     "stuff",
			in3:     "other things",
			addAll:  true,
			prepend: []string{},
			length:  3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			q := gods.NewQueue()

			for _, v := range tC.prepend {
				q.Enqueue(v)
			}

			if tC.addAll {
				q.Enqueue(tC.in1).Enqueue(tC.in2).Enqueue(tC.in3)
			} else {
				q.Enqueue(tC.in1)
			}

			if q.Size() != tC.length {
				t.Errorf("Expected length to be %d but got %d", tC.length, q.Size())
			}
		})
	}
}

func TestQueueDequeue(t *testing.T) {
	testCases := []struct {
		desc        string
		content     []int
		expectError bool
		result      interface{}
	}{
		{
			desc:        "can dequeue from a queue with some stuff",
			content:     []int{1, 2, 3, 4},
			expectError: false,
			result:      1,
		},
		{
			desc:        "can dequeue from a queue with one stuff",
			content:     []int{5},
			expectError: false,
			result:      5,
		},
		{
			desc:        "dequeue from a queue with no stuff returns error",
			content:     []int{},
			expectError: true,
			result:      nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			q := gods.NewQueue()

			for _, v := range tC.content {
				q.Enqueue(v)
			}

			result, err := q.Dequeue()
			if tC.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}

			if result != tC.result {
				t.Errorf("Expected pop to return %d but got %d", tC.result, result)
			}
		})
	}
}

func TestQueuePeek(t *testing.T) {
	testCases := []struct {
		desc        string
		content     []int
		expectError bool
		result      interface{}
	}{
		{
			desc:        "can peek at a queue with some stuff",
			content:     []int{1, 2, 3, 4},
			expectError: false,
			result:      1,
		},
		{
			desc:        "can peek at a queue with one stuff",
			content:     []int{5},
			expectError: false,
			result:      5,
		},
		{
			desc:        "peek at a queue with no stuff returns error",
			content:     []int{},
			expectError: true,
			result:      nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			q := gods.NewQueue()

			for _, v := range tC.content {
				q.Enqueue(v)
			}

			result, err := q.Peek()
			if tC.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}

			if result != tC.result {
				t.Errorf("Expected getAt to return %d but got %d", tC.result, result)
			}
		})
	}
}

func TestQueueIsEmpty(t *testing.T) {
	testCases := []struct {
		desc     string
		content  []int
		expected bool
	}{
		{
			desc:     "returns true for an empty queue",
			content:  []int{},
			expected: true,
		},
		{
			desc:     "returns true for an empty queue",
			content:  []int{1, 2, 3},
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			q := gods.NewQueue()

			for _, v := range tC.content {
				q.Enqueue(v)
			}

			if actual := q.IsEmpty(); actual != tC.expected {
				t.Errorf("Expected IsEmpty to return %t but got %t", tC.expected, actual)
			}
		})
	}
}
