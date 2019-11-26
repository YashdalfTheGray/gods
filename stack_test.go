package gods_test

import (
	"testing"

	"github.com/YashdalfTheGray/gods"
)

func TestNewStack(t *testing.T) {
	testCases := []struct {
		desc string
		out  gods.Stack
	}{
		{
			desc: "Creates a new stack",
			out:  gods.Stack{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if s := gods.NewStack(); s.Size() != 0 {
				t.Errorf("Expected length to be 0 but got %d", s.Size())
			}
		})
	}
}

func TestStackSize(t *testing.T) {
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
			s := gods.NewStack()

			for _, v := range tC.content {
				s.Push(v)
			}

			if listLength := s.Size(); listLength != tC.length {
				t.Errorf("Expected length to be %d but got %d", tC.length, listLength)
			}
		})
	}
}

func TestStackPush(t *testing.T) {
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
			desc:    "can push something on an empty stack",
			in1:     "things",
			in2:     "",
			in3:     "",
			addAll:  false,
			prepend: []string{},
			length:  1,
		},
		{
			desc:    "can push something on a non-empty stack",
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
			s := gods.NewStack()

			for _, v := range tC.prepend {
				s.Push(v)
			}

			if tC.addAll {
				s.Push(tC.in1).Push(tC.in2).Push(tC.in3)
			} else {
				s.Push(tC.in1)
			}

			if s.Size() != tC.length {
				t.Errorf("Expected length to be %d but got %d", tC.length, s.Size())
			}
		})
	}
}

func TestStackPop(t *testing.T) {
	testCases := []struct {
		desc        string
		content     []int
		expectError bool
		result      interface{}
	}{
		{
			desc:        "can pop from a stack with some stuff",
			content:     []int{1, 2, 3, 4},
			expectError: false,
			result:      4,
		},
		{
			desc:        "can pop from a stack with one stuff",
			content:     []int{5},
			expectError: false,
			result:      5,
		},
		{
			desc:        "pop from a stack with no stuff returns error",
			content:     []int{},
			expectError: true,
			result:      nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			s := gods.NewStack()

			for _, v := range tC.content {
				s.Push(v)
			}

			result, err := s.Pop()
			if tC.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}

			if result != tC.result {
				t.Errorf("Expected pop to return %d but got %d", tC.result, result)
			}
		})
	}
}

func TestStackPeek(t *testing.T) {
	testCases := []struct {
		desc        string
		content     []int
		expectError bool
		result      interface{}
	}{
		{
			desc:        "can peek at a stack with some stuff",
			content:     []int{1, 2, 3, 4},
			expectError: false,
			result:      4,
		},
		{
			desc:        "can peek at a stack with one stuff",
			content:     []int{5},
			expectError: false,
			result:      5,
		},
		{
			desc:        "can peek at a stack with no stuff returns error",
			content:     []int{},
			expectError: true,
			result:      nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			s := gods.NewStack()

			for _, v := range tC.content {
				s.Push(v)
			}

			result, err := s.Peek()
			if tC.expectError && err == nil {
				t.Errorf("Expected error but got nil")
			}

			if result != tC.result {
				t.Errorf("Expected getAt to return %d but got %d", tC.result, result)
			}
		})
	}
}

func TestStackIsEmpty(t *testing.T) {
	testCases := []struct {
		desc     string
		content  []int
		expected bool
	}{
		{
			desc:     "returns true for an empty stack",
			content:  []int{},
			expected: true,
		},
		{
			desc:     "returns true for an empty stack",
			content:  []int{1, 2, 3},
			expected: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			s := gods.NewStack()

			for _, v := range tC.content {
				s.Push(v)
			}

			if actual := s.IsEmpty(); actual != tC.expected {
				t.Errorf("Expected IsEmpty to return %t but got %t", tC.expected, actual)
			}
		})
	}
}
