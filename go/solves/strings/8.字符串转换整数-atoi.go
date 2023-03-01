package strings

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	flag := 1
	res := 0
	start := false
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			start = true
			res = res*10 + int(s[i]-'0')
			if flag*res > math.MaxInt32 {
				return math.MaxInt32
			}
			if flag*res < math.MinInt32 {
				return math.MinInt32
			}
		} else {
			if start == true {
				break
			}
			if s[i] == ' ' {
				continue
			} else if s[i] == '+' {
				start = true
				flag = 1
			} else if s[i] == '-' {
				start = true
				flag = -1
			} else {
				break
			}
		}
	}
	res = flag * res
	return res
}

// @lc code=end
