package gods_test

import (
	"testing"

	"github.com/YashdalfTheGray/gods"
)

func TestNewDoublyLinkedList(t *testing.T) {
	testCases := []struct {
		desc     string
		circular bool
	}{
		{
			desc:     "can create a new doubly linked list",
			circular: false,
		},
		{
			desc:     "can create a new doubly linked list",
			circular: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ll := gods.NewDoublyLinkedList(tC.circular)
			if ll.Length() != 0 {
				t.Errorf("Expecting new doubly linked list to exist")
			}
		})
	}
}
