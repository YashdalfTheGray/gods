package gods_test

import (
	"testing"

	"github.com/YashdalfTheGray/gods"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		desc string
		in   interface{}
		out  *gods.ListNode
	}{
		{
			desc: "creates a new list node with the data",
			in:   5,
			out:  &gods.ListNode{Data: 5, Next: nil, Prev: nil},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if ln := gods.NewListNode(tC.in); ln.Data != tC.out.Data {
				t.Errorf("Expected %s but got %s", tC.in, tC.out.Data)
			}
		})
	}
}
