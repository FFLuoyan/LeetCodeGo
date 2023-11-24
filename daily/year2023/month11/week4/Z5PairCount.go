package main

import (
	"sort"
)

func countPairs(nums []int, target int) int {
	result, start, end := 0, 0, len(nums)-1
	sort.Ints(nums)
	for start < end {
		for start < end && nums[start]+nums[end] >= target {
			end--
		}
		result += end - start
		start++
	}
	return result
}
