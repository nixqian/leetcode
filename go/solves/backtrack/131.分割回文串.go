package backtrack

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	res := [][]string{}
	path := []string{}
	var bp func(start int)
	bp = func(start int) {
		if start == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		for i := start; i < len(s); i++ {
			cur := s[start : i+1]
			if !isHuiwen(cur) {
				continue
			}
			path = append(path, cur)
			bp(i + 1)
			path = path[:len(path)-1]
		}
	}
	bp(0)
	return res
}

func isHuiwen(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// @lc code=end
