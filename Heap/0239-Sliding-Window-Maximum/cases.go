package main

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	nums     []int
	k        int
	expected []int
}{
	{[]int{-6, -10, -7, -1, -9, 9, -8, -4, 10, -5, 2, 9, 0, -7, 7, 4, -2, -10, 8, 7}, 7, []int{9, 9, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 8, 8}},
	{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
	{[]int{1, 3, 1, 2, 0, 5}, 3, []int{3, 3, 2, 5}},
	{[]int{}, 3, []int{}},
	{[]int{}, 0, []int{}},
}
