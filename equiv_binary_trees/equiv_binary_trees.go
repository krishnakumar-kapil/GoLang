package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		close(ch)
		return
	}
	helper(t, ch)
}

func helper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	helper(t.Left, ch)
	ch <- t.Value
	helper(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
}
