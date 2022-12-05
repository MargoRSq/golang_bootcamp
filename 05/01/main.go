package main

import "fmt"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func addRightToLeft(root *TreeNode, level int, direction bool) (snake []bool) {
	if root == nil {
		return
	}

	if level == 1 {
		snake = append(snake, root.HasToy)
	} else if level > 1 && direction {
		snake = append(snake, addRightToLeft(root.Left, level-1, direction)...)
		snake = append(snake, addRightToLeft(root.Right, level-1, direction)...)
	} else if level > 1 && !direction {
		snake = append(snake, addRightToLeft(root.Right, level-1, direction)...)
		snake = append(snake, addRightToLeft(root.Left, level-1, direction)...)
	}
	return
}

func unrollGarland(root *TreeNode) (snake []bool) {
	flag := false
	height := root.maxDepth()

	for i := 1; i <= height; i++ {
		if flag {
			snake = append(snake, addRightToLeft(root, i, flag)...)
			flag = false
		} else {
			snake = append(snake, addRightToLeft(root, i, flag)...)
			flag = true
		}
	}
	return
}

func main() {

	root := &TreeNode{HasToy: false}
	root.Left = &TreeNode{HasToy: false}
	root.Left.Left = &TreeNode{HasToy: false}
	root.Left.Right = &TreeNode{HasToy: true}
	root.Right = &TreeNode{HasToy: true}
	root.Right.Right = &TreeNode{HasToy: true}

	snake := unrollGarland(root)
	fmt.Println(snake)

}
