package main

import (
	"reflect"
	"testing"
)

func createTests() (Trees []*TreeNode) {
	Trees = make([]*TreeNode, 9)
	Trees[0] = &TreeNode{false,
		&TreeNode{false,
			&TreeNode{false, nil, nil},
			&TreeNode{true, nil, nil}},
		&TreeNode{true, nil, nil}}
	Trees[1] = &TreeNode{true,
		&TreeNode{true,
			&TreeNode{true, nil, nil},
			&TreeNode{false, nil, nil}},
		&TreeNode{false,
			&TreeNode{true, nil, nil},
			&TreeNode{true, nil, nil}}}
	Trees[2] = &TreeNode{true,
		&TreeNode{true, nil, nil},
		&TreeNode{false, nil, nil}}
	Trees[3] = &TreeNode{false,
		&TreeNode{true, nil,
			&TreeNode{true, nil, nil}},
		&TreeNode{false, nil,
			&TreeNode{true, nil, nil}}}
	Trees[4] = &TreeNode{true, nil, nil}
	Trees[5] = &TreeNode{false, nil, nil}
	Trees[6] = &TreeNode{false, nil, &TreeNode{false, nil, nil}}
	Trees[7] = &TreeNode{false, nil, &TreeNode{true, nil, nil}}
	Trees[8] = &TreeNode{true,
		&TreeNode{true,
			&TreeNode{true, nil, &TreeNode{true,
				&TreeNode{true,
					&TreeNode{true, nil, nil},
					&TreeNode{false, nil, nil}},
				&TreeNode{false,
					&TreeNode{true, nil, nil},
					&TreeNode{true, nil, nil}}}},
			&TreeNode{false, nil, nil}},
		&TreeNode{false,
			&TreeNode{true, nil, nil},
			&TreeNode{true, nil, nil}}}
	return
}

var trees []*TreeNode = createTests()
var tests = []struct {
	input  *TreeNode
	output []bool
}{
	{trees[0], []bool{false, false, true, true, false}},
	{trees[1], []bool{true, true, false, true, true, false, true}},
	{trees[2], []bool{true, true, false}},
	{trees[3], []bool{false, true, false, true, true}},
	{trees[4], []bool{true}},
	{trees[5], []bool{false}},
	{trees[6], []bool{false, false}},
	{trees[7], []bool{false, true}},
	{trees[8], []bool{true, true, false, true, true, false, true, true, false, true, true, false, true, true}},
}

func TestBalanced(t *testing.T) {
	for i, test := range tests {
		if res := unrollGarland(test.input); !reflect.DeepEqual(res, test.output) {
			t.Errorf("%d test failed\n exp: %v\n got: %v", i, res, test.output)
		}
	}
}
