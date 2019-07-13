package main

import "math"

/*******************************************************************************
  TestTable
*******************************************************************************/

var TestCases = []struct {
	str      string
	expected int
}{
	{"42", 42},
	{"   -42", -42},
	{"+0", 0},
	{"+", 0},
	{"-", 0},
	{"+w", 0},
	{"", 0},
	{"1000005", 1000005},
	{"105", 105},
	{"-w", 0},
	{"4193 with words", 4193},
	{"words and 987", 0},

	// floats
	{"2", 2},
	{"2.0", 2},
	{"2.1", 2},
	{"-2.1", -2},

	{"010", 10},

	// overflow
	{"-91283472332", math.MinInt32},
	{"9223372036854775808", math.MaxInt32},

	{"2147483647", math.MaxInt32},
	{"2147483646", 2147483646},
	{"21474836471", math.MaxInt32},
	{"-2147483648", math.MinInt32},
	{"-2147483647", -2147483647},
	{"-21474836471", math.MinInt32},

	{"1000000000000000000000000000000000000000000522545459", math.MaxInt32},
	{"-100000000000000000000000000000000000000000522545459", math.MinInt32},
	{"10000000000000000000000000000000000000000000000000000000522545459", math.MaxInt32},
	{"-10000000000000000000000000000000000000000000000000000000522545459", math.MinInt32},

	{"  0000000000012345678", 12345678},
	{"012345678", 12345678},
	{"0-1", 0},
	{"20000000000000000000", math.MaxInt32},
	{"-6147483648", math.MinInt32},
	{"      -11919730356x", math.MinInt32},
	{"   +0 123", 0},
	{"00214748364000007", math.MaxInt32},
	{"000021474836473432398249801273402189347901283470918234790183247901203984", math.MaxInt32},
	{"     +004500", 4500},
	{"     -004500", -4500},
	{"     +0040500", 40500},
	{"     -0040500", -40500},
}
