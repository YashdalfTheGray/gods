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
