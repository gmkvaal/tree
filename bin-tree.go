package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	if t.Right != nil {
		ch <- t.Right.Value
		Walk(t.Right, ch)
	}

	if t.Left != nil {
		ch <- t.Left.Value
		Walk(t.Left, ch)
	}
}




// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool {

//}

func main() {

	fmt.Println(tree.New(1))

	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)


	for i := range ch {
		fmt.Println(i)
	}

}


