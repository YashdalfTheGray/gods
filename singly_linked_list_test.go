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
		length  uint32
	}{
		{
			desc:    "can put something in an empty list",
			in1:     "things",
			in2:     "",
			in3:     "",
			addAll:  false,
			prepend: []string{},
			length:  1,
		},
		{
			desc:    "can put something in a non-empty list",
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
			ll := gods.NewSinglyLinkedList()

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

func TestSinglyLinkedListUnshift(t *testing.T) {
	testCases := []struct {
		desc        string
		content     []int
		expectError bool
		result      interface{}
	}{
		{
			desc:        "can unshift from a list with some stuff",
			content:     []int{1, 2, 3, 4},
			expectError: false,
			result:      1,
		},
		{
			desc:        "can unshift from a list with one stuff",
			content:     []int{5},
			expectError: false,
			result:      5,
		},
		{
			desc:        "unshift from a list with no stuff returns error",
			content:     []int{},
			expectError: true,
			result:      nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ll := gods.NewSinglyLinkedList()

			for i := len(tC.content) - 1; i >= 0; i-- {
				ll.Shift(tC.content[i])
			}

			result, err := ll.Unshift()
			if tC.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}

			if result != tC.result {
				t.Errorf("Expected unshift to return %d but got %d", tC.result, result)
			}
		})
	}
}

func TestSinglyLinkedListPush(t *testing.T) {
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
			desc:    "can put something in an empty list",
			in1:     "things",
			in2:     "",
			in3:     "",
			addAll:  false,
			prepend: []string{},
			length:  1,
		},
		{
			desc:    "can put something in a non-empty list",
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
			ll := gods.NewSinglyLinkedList()

			for _, v := range tC.prepend {
				ll.Push(v)
			}

			if tC.addAll {
				ll.Push(tC.in1).Push(tC.in2).Push(tC.in3)
			} else {
				ll.Push(tC.in1)
			}

			if ll.Length() != tC.length {
				t.Errorf("Expected length to be %d but got %d", tC.length, ll.Length())
			}
		})
	}
}

func TestSinglyLinkedListPop(t *testing.T) {
	testCases := []struct {
		desc        string
		content     []int
		expectError bool
		result      interface{}
	}{
		{
			desc:        "can pop from a list with some stuff",
			content:     []int{1, 2, 3, 4},
			expectError: false,
			result:      4,
		},
		{
			desc:        "can pop from a list with one stuff",
			content:     []int{5},
			expectError: false,
			result:      5,
		},
		{
			desc:        "pop from a list with no stuff returns error",
			content:     []int{},
			expectError: true,
			result:      nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ll := gods.NewSinglyLinkedList()

			for _, v := range tC.content {
				ll.Push(v)
			}

			result, err := ll.Pop()
			if tC.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}

			if result != tC.result {
				t.Errorf("Expected unshift to return %d but got %d", tC.result, result)
			}
		})
	}
}

func TestSinglyLinkedListIterator(t *testing.T) {
	testCases := []struct {
		desc    string
		content []int
	}{
		{
			desc:    "iterates over a linked list with some stuff",
			content: []int{1, 2, 3, 4},
		},
		{
			desc:    "iterates over a linked list with one stuff",
			content: []int{5},
		},
		{
			desc:    "iterates over a linked list with no stuff",
			content: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ll := gods.NewSinglyLinkedList()

			for i := len(tC.content) - 1; i >= 0; i-- {
				ll.Shift(tC.content[i])
			}

			iterator := ll.Iterate()
			for i := 0; iterator.Next(); {
				if v := iterator.Get(); v != tC.content[i] {
					t.Errorf("Expected %d at position %d but got %d", tC.content[i], i, v)
				}
				i++
			}

			if err := iterator.Error(); err != nil {
				t.Errorf("Expected the iterator to not be in an error state")
			}
		})
	}
}
