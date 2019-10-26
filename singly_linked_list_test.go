package gods_test

import (
	"testing"

	"github.com/YashdalfTheGray/gods"
)

func TestNewSinglyLinkedList(t *testing.T) {
	testCases := []struct {
		desc string
		out  gods.SinglyLinkedList
	}{
		{
			desc: "Creates a new singly linked list",
			out:  gods.SinglyLinkedList{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if ll := gods.NewSinglyLinkedList(); ll.Length() != 0 {
				t.Errorf("Expected length to be 0 but got %d", ll.Length())
			}
		})
	}
}
