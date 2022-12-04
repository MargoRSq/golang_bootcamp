package main

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func (root *TreeNode) countValues(value int) (result int) {
	if root == nil {
		return
	}
	result += root.Left.countValues(1)
	if root.HasToy {
		result += 1
	}
	result += root.Right.countValues(1)
	return
}

func areToysBalanced(root *TreeNode) (isBalanced bool) {
	leftNode, rightNode := root.Left, root.Right
	if left, right := leftNode.countValues(1), rightNode.countValues(1); left != right {
		isBalanced = false
	} else {
		isBalanced = true
	}
	return
}

// func main() {

// 	root := &TreeNode{HasToy: false}
// 	root.Left = &TreeNode{HasToy: false}
// 	root.Left.Left = &TreeNode{HasToy: false}
// 	root.Left.Right = &TreeNode{HasToy: true}
// 	root.Right = &TreeNode{HasToy: true}
// 	root.Right.Right = &TreeNode{HasToy: true}

// 	isBalanced := areToysBalanced(root)
// 	fmt.Println(isBalanced)

// }
