package maths

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	flag := 1
	if x < 0 {
		flag = -1
		x = x * -1
	}
	res := 0
	for x > 0 {
		res = res*10 + x%10
		x = x / 10
	}
	res = res * flag
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}

// @lc code=end
