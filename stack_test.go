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
