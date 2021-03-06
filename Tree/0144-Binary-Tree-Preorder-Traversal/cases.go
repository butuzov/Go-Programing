package main

import (
	"strings"

	"github.com/butuzov/leetcode.go/pkg/binarytree"
)

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	root     *binarytree.TreeNode
	expected []int
}{
	{
		binarytree.NewTree(strings.Split("25|12l|10l|4l|1l|p|p|p|5r|p|p|4r|0r|31r|100r", "|")),
		[]int{25, 12, 10, 4, 1, 5, 4, 0, 31, 100},
	},
	{
		binarytree.NewTree(strings.Split("25|12l|6l|p|17r|p|p|37r|29l|p|50r", "|")),
		[]int{25, 12, 6, 17, 37, 29, 50},
	},
	{
		binarytree.NewTree(strings.Split("1|2r|3l", "|")),
		[]int{1, 2, 3},
	},
}
