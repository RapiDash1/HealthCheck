package bst

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	for i := 0; i < 100; i++ {
		AddNode(i)
	}
	fmt.Println(SortedOrder())
}
