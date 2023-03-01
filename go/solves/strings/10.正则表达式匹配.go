package strings

import "regexp"

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	reg := regexp.MustCompile(p)
	return reg.MatchString(s)
}

// @lc code=end
