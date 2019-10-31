package gods_test

import (
	"fmt"

	"github.com/YashdalfTheGray/gods"
)

func ExampleIterator() {
	ll := gods.NewSinglyLinkedList()

	ll.Push(1).Push(2).Push(3).Push(4)

	iterator := ll.Iterate()
	for iterator.Next() {
		fmt.Println(iterator.Get())
	}

	if err := iterator.Error(); err != nil {
		fmt.Println(err.Error())
	}
}
