package problems

// https://leetcode.com/problems/two-sum/

// TwoSum returns indexes of two elements which sum is equal to a target
func TwoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
