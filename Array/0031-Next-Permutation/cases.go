package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	nums     []int
	expected []int
}{
	{[]int{}, []int{}},
	{[]int{1, 1}, []int{1, 1}},
	{[]int{1, 2, 3}, []int{1, 3, 2}},
	{[]int{3, 2, 1}, []int{1, 2, 3}},
	{[]int{1, 1, 5}, []int{1, 5, 1}},
	{[]int{1, 3, 2}, []int{2, 1, 3}},
	{[]int{2, 3, 1}, []int{3, 1, 2}},
	{[]int{1, 2, 3, 2, 1}, []int{1, 3, 1, 2, 2}},
	{[]int{5, 4, 7, 5, 3, 2}, []int{5, 5, 2, 3, 4, 7}},
	{[]int{4, 2, 0, 2, 3, 2, 0}, []int{4, 2, 0, 3, 0, 2, 2}},
	{[]int{1, 2, 5, 4, 3, 2, 1}, []int{1, 3, 1, 2, 2, 4, 5}},
	{[]int{2, 2, 7, 5, 4, 3, 2, 2, 1}, []int{2, 3, 1, 2, 2, 2, 4, 5, 7}},
}