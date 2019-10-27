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

func TestSinglyLinkedListShift(t *testing.T) {
	testCases := []struct {
		desc    string
		in1     string
		in2     string
		in3     string
		addAll  bool
		prepend []string
		legnth  uint32
	}{
		{
			desc:    "can put something in an empty list",
			in1:     "things",
			in2:     "",
			in3:     "",
			addAll:  false,
			prepend: []string{},
			legnth:  1,
		},
		{
			desc:    "can put something in a non-empty list",
			in1:     "bar",
			in2:     "",
			in3:     "",
			addAll:  false,
			prepend: []string{"foo"},
			legnth:  2,
		},
		{
			desc:    "supports fluent API",
			in1:     "things",
			in2:     "stuff",
			in3:     "other things",
			addAll:  true,
			prepend: []string{},
			legnth:  3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ll := gods.NewSinglyLinkedList()

			for _, v := range tC.prepend {
				ll.Shift(v)
			}

			if tC.addAll {
				ll.Shift(tC.in1).Shift(tC.in2).Shift(tC.in3)
			} else {
				ll.Shift(tC.in1)
			}

			if ll.Length() != tC.legnth {
				t.Errorf("Expected length to be 1 but got %d", ll.Length())
			}
		})
	}
}
