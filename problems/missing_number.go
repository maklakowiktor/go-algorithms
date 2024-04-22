package problems

// LeetCode: 268. Missing Number
// https://leetcode.com/problems/missing-number

func MissingNumber(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	sumAll := (len(nums) * (len(nums) + 1)) / 2
	return sumAll - sum
}
