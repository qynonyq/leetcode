package problems

// https://leetcode.com/problems/two-sum/

// TwoSum returns indexes of two elements which sum is equal to a target
func TwoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	for i, n := range nums {
		for j := i + 1; j < len(nums); j++ {
			if n+nums[j] == target {
				result = []int{i, j}
				break
			}
		}
	}
	return result
}
