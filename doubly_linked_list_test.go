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

func TestDoublyLinkedListLength(t *testing.T) {
	testCases := []struct {
		desc    string
		content []int
		length  uint32
	}{
		{
			desc:    "returns 0 for empty linked list",
			content: []int{},
			length:  0,
		},
		{
			desc:    "returns 1 for a linked list with 1 item in it",
			content: []int{1},
			length:  1,
		},
		{
			desc:    "returns length of linked list",
			content: []int{1, 2, 3, 4},
			length:  4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ll := gods.NewSinglyLinkedList()

			for _, v := range tC.content {
				ll.Shift(v)
			}

			if listLength := ll.Length(); listLength != tC.length {
				t.Errorf("Expected length to be %d but got %d", tC.length, listLength)
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

func TestDoublyLinkedListUnshift(t *testing.T) {
	testCases := []struct {
		desc        string
		content     []int
		expectError bool
		circular    bool
		result      interface{}
	}{
		{
			desc:        "can unshift from a list with some stuff",
			content:     []int{1, 2, 3, 4},
			expectError: false,
			circular:    false,
			result:      1,
		},
		{
			desc:        "can unshift from a list with one stuff",
			content:     []int{5},
			expectError: false,
			circular:    false,
			result:      5,
		},
		{
			desc:        "unshift from a list with no stuff returns error",
			content:     []int{},
			expectError: true,
			circular:    false,
			result:      nil,
		},
		{
			desc:        "can unshift from a circular list with some stuff",
			content:     []int{1, 2, 3, 4},
			expectError: false,
			circular:    true,
			result:      1,
		},
		{
			desc:        "can unshift from a circular list with one stuff",
			content:     []int{5},
			expectError: false,
			circular:    true,
			result:      5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ll := gods.NewDoublyLinkedList(tC.circular)

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

func TestDoublyLinkedListPush(t *testing.T) {
	testCases := []struct {
		desc     string
		in1      string
		in2      string
		in3      string
		addAll   bool
		prepend  []string
		circular bool
		length   uint32
	}{
		{
			desc:     "can put something in an empty list",
			in1:      "things",
			in2:      "",
			in3:      "",
			addAll:   false,
			prepend:  []string{},
			circular: false,
			length:   1,
		},
		{
			desc:     "can put something in a non-empty list",
			in1:      "bar",
			in2:      "",
			in3:      "",
			addAll:   false,
			prepend:  []string{"foo"},
			circular: false,
			length:   2,
		},
		{
			desc:     "supports fluent API",
			in1:      "things",
			in2:      "stuff",
			in3:      "other things",
			addAll:   true,
			prepend:  []string{},
			circular: false,
			length:   3,
		},
		{
			desc:     "can put something in an empty circular list",
			in1:      "things",
			in2:      "",
			in3:      "",
			addAll:   false,
			prepend:  []string{},
			circular: true,
			length:   1,
		},
		{
			desc:     "can put something in a non-empty circular list",
			in1:      "bar",
			in2:      "",
			in3:      "",
			addAll:   false,
			prepend:  []string{"foo"},
			circular: true,
			length:   2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ll := gods.NewDoublyLinkedList(tC.circular)

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

func TestDoublyLinkedListPop(t *testing.T) {
	testCases := []struct {
		desc        string
		content     []int
		circular    bool
		expectError bool
		result      interface{}
	}{
		{
			desc:        "can pop from a list with some stuff",
			content:     []int{1, 2, 3, 4},
			circular:    false,
			expectError: false,
			result:      4,
		},
		{
			desc:        "can pop from a list with one stuff",
			content:     []int{5},
			circular:    false,
			expectError: false,
			result:      5,
		},
		{
			desc:        "pop from a list with no stuff returns error",
			content:     []int{},
			circular:    false,
			expectError: true,
			result:      nil,
		},
		{
			desc:        "can pop from a circular list with some stuff",
			content:     []int{1, 2, 3, 4},
			circular:    true,
			expectError: false,
			result:      4,
		},
		{
			desc:        "can pop from a circular list with one stuff",
			content:     []int{5},
			circular:    true,
			expectError: false,
			result:      5,
		},
		{
			desc:        "pop from a circular list with no stuff returns error",
			content:     []int{},
			circular:    true,
			expectError: true,
			result:      nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ll := gods.NewDoublyLinkedList(tC.circular)

			for _, v := range tC.content {
				ll.Push(v)
			}

			result, err := ll.Pop()
			if tC.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}

			if result != tC.result {
				t.Errorf("Expected pop to return %d but got %d", tC.result, result)
			}
		})
	}
}
