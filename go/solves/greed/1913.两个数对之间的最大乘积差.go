package greed

import "sort"

/*
 * @lc app=leetcode.cn id=1913 lang=golang
 *
 * [1913] 两个数对之间的最大乘积差
 */

// @lc code=start
func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]*nums[len(nums)-2] - nums[0]*nums[1]
}

// @lc code=end
