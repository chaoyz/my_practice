package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func main() {
	var tree1 Tree = buildTree1()
	var tree2 Tree = buildTree2()
	fmt.Println(tree1.Value)
	fmt.Println(tree2.Value)
	// array1 := make([]int, 10)
	// array2 := make([]int, 10)
	getTreeArray(tree1)
	// getTreeArray(*tree2, array2)
	// fmt.Println(array1)
	// fmt.Println(array2)
}

func getTreeArray(tree Tree) {
	if tree.Left.Value == 0 && tree.Right.Value == 0 {
		// append(array, *tree.Value)
		fmt.Println(tree.Value)
		return
	} else {
		if tree.Left.Value != 0 {
			// append(array, *tree.Left.Value)
			fmt.Println(tree.Value)
			fmt.Println(tree.Left)
			fmt.Println(tree.Left == (*Tree)(nil))
			if tree.Left != (*Tree)(nil) {
				getTreeArray(*tree.Left)
			}
		}
		// if tree.Right.Value != 0 {
		// 	// append(array, *tree.Right.Value)
		// 	fmt.Println(tree.Value)
		// 	getTreeArray(*tree.Right, array)
		// }
	}
}

func buildTree1() Tree {
	var root Tree
	left := Tree{&Tree{Value: 1}, 1, &Tree{Value: 2}}
	right := Tree{&Tree{Value: 5}, 8, &Tree{Value: 13}}
	root.Left = &left
	root.Right = &right
	root.Value = 3
	return root
}

func buildTree2() Tree {
	var root Tree
	left := Tree{&Tree{Value: 1}, 1, &Tree{Value: 2}}
	l := Tree{&left, 3, &Tree{Value: 5}}
	root.Left = &l
	root.Right = &Tree{Value: 13}
	root.Value = 8
	return root
}
