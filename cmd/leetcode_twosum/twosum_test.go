package main

import (
	"fmt"

	"github.com/sabrusrin/wildberries_st5/pkg/leetcode_twosum"
)

func TestOk(t *testing.T)	{
	testcase := []int{3, 2, 4}
	out := leetcode_twosum.TwoSum(testcase, 6)
	if out[0] != 1 || out[1] != 2	{
		t.Errorf("Incorrect indexes!\nExpected: %v\tGot: %v\n", "[1 2]", out)
	}
	if (testcase[out[0]] + testcase[out[1]]) != 6	{
		t.Errorf("Incorrect sum!\nExpected: %v\tGot: %v\n", 6, testcase[out[0]] + testcase[out[1]])
	}
	testcase = []int{2, 7, 11, 15}
	out = leetcode_twosum.TwoSum(testcase, 9)
	if out[0] != 0 || out[1] != 1	{
		t.Errorf("Incorrect indexes!\nExpected: %v\tGot: %v\n", "[0 1]", out)
	}
	if (testcase[out[0]] + testcase[out[1]]) != 9	{
		t.Errorf("Incorrect sum!\nExpected: %v\tGot: %v\n", 6, testcase[out[0]] + testcase[out[1]])
	}
}
