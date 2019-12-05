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
			tn := gods.NewTreeNode(tC.data)

			if len(tn.Children) != 0 {
				t.Errorf("Expected initial children length to be 0")
			}

			if tn.Data != tC.data {
				t.Errorf("Expected data to be %d but got %d", tC.data, tn.Data)
			}
		})
	}
}

func TestTreeNodeIsLeaf(t *testing.T) {
	testCases := []struct {
		desc     string
		data     int
		children []int
	}{
		{
			desc:     "returns true for a leaf node",
			data:     1,
			children: []int{},
		},
		{
			desc:     "returns false for an actual tree node",
			data:     1,
			children: []int{1, 2, 3},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tn := gods.NewTreeNode(tC.data)

			for _, v := range tC.children {
				tn.AddChild(v)
			}

			if tn.IsLeaf() != (len(tC.children) == 0) {
				t.Errorf("Expected tree node to be a leaf and it wasn't")
			}
		})
	}
}

func TestTreeNodeAddChild(t *testing.T) {
	testCases := []struct {
		desc            string
		child           uint8
		prepend         []uint8
		numberOfChilden int
	}{
		{
			desc:            "can add a child to the tree node",
			child:           8,
			prepend:         []uint8{},
			numberOfChilden: 1,
		},
		{
			desc:            "can add a child to the tree node with children",
			child:           8,
			prepend:         []uint8{2, 3, 4},
			numberOfChilden: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			tn := gods.NewTreeNode(1)

			for _, childToAdd := range tC.prepend {
				tn.AddChild(childToAdd)
			}
			tn.AddChild(tC.child)

			if len(tn.Children) != tC.numberOfChilden {
				t.Errorf("Expected children to have length %d but got %d", tC.numberOfChilden, len(tn.Children))
			}
		})
	}
}
