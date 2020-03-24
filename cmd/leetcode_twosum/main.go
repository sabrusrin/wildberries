package main

import (
	"fmt"

	"github.com/sabrusrin/wildberries_st5/pkg/leetcode_twosum"
)

func main()	{
	nums := []int{2, 7, 11, 15}
	out := leetcode_twosum.TwoSum(nums, 9)
	fmt.Printf("Indices of the two numbers such that they add up to 9: %v\n", out)
}
