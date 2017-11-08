package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	// close(ch)
}

func walk(t *tree.Tree, ch chan int) {

	if t != nil {
		ch <- t.Value
		go Walk(t.Left, ch)
		go Walk(t.Right, ch)
	}
}

// SameValuesInSlice hecks if the equal length
// slices s1 and s2 contains the same values
func SameValuesInSlice(s1, s2 []int) bool {

	counter := 0
	for _, nodeVal1 := range s1 {
		for _, nodeVal2 := range s2 {
			if nodeVal1 == nodeVal2 {
				counter += 1
			}
		}
	}

	if counter == len(s1) {
		return true
	}

	return false
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	var nodeVals1 []int
	var nodeVals2 []int

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		nodeVals1 = append(nodeVals1, <-ch1)
		nodeVals2 = append(nodeVals2, <-ch2)
	}

	return SameValuesInSlice(nodeVals1, nodeVals2)
}

func main() {

	fmt.Println(Same(tree.New(1), tree.New(1)))

	/*

	ch := make(chan int, 10)
	var nodeVals []int

	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		nodeVals = append(nodeVals, <-ch)
	}

	fmt.Println(nodeVals)
	*/



}


