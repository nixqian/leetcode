package maths

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	reverse := 0
	tmp := x
	for tmp > 0 {
		reverse = reverse*10 + tmp%10
		tmp = tmp / 10
	}
	return x == reverse
}

// @lc code=end
