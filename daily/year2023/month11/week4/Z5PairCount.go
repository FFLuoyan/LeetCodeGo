package main

func countPairs(nums []int, target int) int {
	result := 0
	for i, left := range nums {
		for _, right := range nums[i+1:] {
			if left+right < target {
				result++
			}
		}
	}
	return result
}
