package array

import "fmt"

/*
 * @lc app=leetcode.cn id=1016 lang=golang
 *
 * [1016] 子串能表示从 1 到 N 数字的二进制串
 */

// @lc code=start
func queryString(s string, n int) bool {
	hasMap := map[string]bool{}
	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			hasMap[s[i:j]] = true
		}
	}
	fmt.Printf("%+v\n", hasMap)
	for i := 1; i <= n; i++ {
		if !hasMap[getBi(i)] {
			return false
		}
	}
	return true
}

func getBi(x int) string {
	fmt.Printf("%b\n", x)
	return fmt.Sprintf("%b", x)
}

// @lc code=end
