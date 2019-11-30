package gods_test

import (
	"testing"

	"github.com/YashdalfTheGray/gods"
)

func TestNewTreeNode(t *testing.T) {
	testCases := []struct {
		desc string
		data interface{}
	}{
		{
			desc: "creates a new tree node with some data",
			data: 5,
		},
		{
			desc: "creates a new tree node with other data",
			data: 10,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tn := gods.NewTreeNode(tC.desc)

			if len(tn.Children) != 0 {
				t.Errorf("Expected initial children length to be 0")
			}
		})
	}
}