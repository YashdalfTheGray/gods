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
			if s := gods.NewStack(); s.Length() != 0 {
				t.Errorf("Expected length to be 0 but got %d", s.Length())
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

			if s.Length() != tC.length {
				t.Errorf("Expected length to be %d but got %d", tC.length, s.Length())
			}
		})
	}
}
