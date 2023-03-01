package bisearch

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
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if divisor == 0 {
		return 0
	}
	flag := 1
	if dividend < 0 {
		flag *= -1
		dividend = 0 - dividend
	}
	if divisor < 0 {
		flag *= -1
		divisor = 0 - divisor
	}
	return flag * fastDivide(dividend, divisor)
}
func fastDivide(a, b int) int {
	step := 1
	dividend := a
	divisor := b
	res := 0
	// 预处理生成最大的除数
	for dividend >= divisor<<1 {
		divisor = divisor << 1
		step = step << 1
	}
	for dividend >= b {
		if dividend >= divisor {
			dividend = dividend - divisor
			res += step
		} else {
			divisor >>= 1
			step >>= 1
		}
	}
	return res
}

// @lc code=end
