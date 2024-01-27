package problems

// LeetCode: 268. Missing Number
// https://leetcode.com/problems/missing-number

func MissingNumber(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return len(nums)*(len(nums)+1)/2 - result
}
