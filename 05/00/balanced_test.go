package main

import "testing"

func createTests() (first *TreeNode, second *TreeNode, third *TreeNode, fourth *TreeNode) {
	first = &TreeNode{HasToy: false}
	first.Left = &TreeNode{HasToy: false}
	first.Left.Left = &TreeNode{HasToy: false}
	first.Left.Right = &TreeNode{HasToy: true}
	first.Right = &TreeNode{HasToy: true}

	second = &TreeNode{HasToy: true}
	second.Left = &TreeNode{HasToy: true}
	second.Left.Left = &TreeNode{HasToy: true}
	second.Left.Right = &TreeNode{HasToy: false}
	second.Right = &TreeNode{HasToy: false}
	second.Right.Left = &TreeNode{HasToy: true}
	second.Right.Right = &TreeNode{HasToy: true}

	third = &TreeNode{HasToy: false}
	third.Left = &TreeNode{HasToy: true}
	third.Right = &TreeNode{HasToy: false}

	fourth = &TreeNode{HasToy: false}
	fourth.Left = &TreeNode{HasToy: true}
	fourth.Left.Right = &TreeNode{HasToy: true}
	fourth.Right = &TreeNode{HasToy: false}
	fourth.Right.Right = &TreeNode{HasToy: true}
	return
}

var first, second, third, fourth *TreeNode = createTests()
var tests = []struct {
	input  *TreeNode
	output bool
}{
	{first, true},
	{second, true},
	{third, false},
	{fourth, false},
}

func TestBalanced(t *testing.T) {
	for i, test := range tests {
		if res := areToysBalanced(test.input); res != test.output {
			t.Errorf("%d test failed", i)
		}
	}
}
