package z5

import (
	"sort"
)

/**
 * countPairs
 * 给定一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 target
 * 请返回满足 0 <= i < j < n 且 nums[i] + nums[j] < target 的下标对 (i, j) 的数目
 *
 * 1 <= nums.length == n <= 50
 * -50 <= nums[i], target <= 50
 *
 * @author Li.zongjie
 * @version 1.0
 * @date 2023/11/24
 */
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
