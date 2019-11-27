package gods_test

import (
	"testing"

	"github.com/YashdalfTheGray/gods"
)

func TestNewQueue(t *testing.T) {
	testCases := []struct {
		desc string
	}{
		{
			desc: "Creates a new queue",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if q := gods.NewQueue(); q.Size() != 0 {
				t.Errorf("Expected length to be 0 but got %d", q.Size())
			}
		})
	}
}
