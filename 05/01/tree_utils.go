package main

import "fmt"

func printTree(prefix string, tree *TreeNode, isLeft bool, level int) {
	if tree != nil {
		fmt.Print(level, prefix)
		if isLeft {
			fmt.Print("├──")
		} else {
			fmt.Print("└──")
		}
		if tree.HasToy {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
		if isLeft {
			printTree(prefix+"│   ", tree.Left, true, level+1)
		} else {
			printTree(prefix+"    ", tree.Left, true, level+1)
		}
		if isLeft {
			printTree(prefix+"│   ", tree.Right, false, level+1)
		} else {
			printTree(prefix+"    ", tree.Right, false, level+1)
		}
	}
}

func (root *TreeNode) maxDepth() int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	lHeight := root.Left.maxDepth()
	rHeight := root.Right.maxDepth()

	if lHeight >= rHeight {
		return lHeight + 1
	} else {
		return rHeight + 1
	}
}
