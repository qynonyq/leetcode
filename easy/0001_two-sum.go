package problems

import "math"

// https://leetcode.com/problems/two-sum/

func TwoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	max := int(math.Pow10(9))
	min := -int(math.Pow10(9))
	if len(nums) < 2 || len(nums) > 10000 {
		return nil
	}
	if target < min || target > max {
		return nil
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > max || nums[i] > max {
			return nil
		}
		if nums[i-1] < min || nums[i] < min {
			return nil
		}
		if nums[i-1]+nums[i] == target {
			result = []int{i - 1, i}
			break
		}
	}
	return result
}
