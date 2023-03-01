package maths

import (
	"math"
	"sort"
)

/*
* @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
*/

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	res := math.MaxInt
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(sum-target)) < math.Abs(float64(res-target)) {
				res = sum
			}
			if sum > target {
				r--
			} else {
				l++
			}
		}
	}
	return res
}

// @lc code=end
