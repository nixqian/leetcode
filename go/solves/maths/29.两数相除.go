package maths

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	// 考虑溢出的情况
	if dividend == math.MinInt32 {
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == 0 {
		return 0
	}
	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	flag := true
	if dividend < 0 {
		flag = !flag
		dividend = 0 - dividend
	}
	if divisor < 0 {
		flag = !flag
		divisor = 0 - divisor
	}
	res := 0
	for dividend >= divisor {
		res += 1
		dividend = dividend - divisor
	}
	if !flag {
		res = 0 - res
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	if res < math.MinInt32 {
		return math.MinInt32
	}
	return res
}

// @lc code=end
