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

func TestDoublyLinkedListShift(t *testing.T) {
	testCases := []struct {
		desc     string
		in1      string
		in2      string
		in3      string
		addAll   bool
		circular bool
		prepend  []string
		length   uint32
	}{
		{
			desc:     "can put something in an empty list",
			in1:      "things",
			in2:      "",
			in3:      "",
			addAll:   false,
			circular: false,
			prepend:  []string{},
			length:   1,
		},
		{
			desc:     "can put something in a non-empty list",
			in1:      "bar",
			in2:      "",
			in3:      "",
			addAll:   false,
			circular: false,
			prepend:  []string{"foo"},
			length:   2,
		},
		{
			desc:     "supports fluent API",
			in1:      "things",
			in2:      "stuff",
			in3:      "other things",
			addAll:   true,
			circular: false,
			prepend:  []string{},
			length:   3,
		},
		{
			desc:     "can put something in an empty circular list",
			in1:      "things",
			in2:      "",
			in3:      "",
			addAll:   false,
			circular: true,
			prepend:  []string{},
			length:   1,
		},
		{
			desc:     "can put something in a non-empty circular list",
			in1:      "bar",
			in2:      "",
			in3:      "",
			addAll:   false,
			circular: true,
			prepend:  []string{"foo"},
			length:   2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ll := gods.NewDoublyLinkedList(tC.circular)

			for _, v := range tC.prepend {
				ll.Shift(v)
			}

			if tC.addAll {
				ll.Shift(tC.in1).Shift(tC.in2).Shift(tC.in3)
			} else {
				ll.Shift(tC.in1)
			}

			if ll.Length() != tC.length {
				t.Errorf("Expected length to be %d but got %d", tC.length, ll.Length())
			}
		})
	}
}
