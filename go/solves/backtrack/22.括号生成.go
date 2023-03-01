package backtrack

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}
	path := ""
	st := []string{}
	var bp func(left int)
	bp = func(left int) {
		//fmt.Println(path)
		if len(path) == n*2 {
			res = append(res, path)
			return
		}
		// 括号入栈
		if left > 0 {
			st = append(st, "(")
			path += "("
			bp(left - 1)
			path = path[:len(path)-1]
			st = st[:len(st)-1]
		}
		// 括号出栈
		if len(st) > 0 {
			st = st[:len(st)-1]
			path += ")"
			bp(left)
			path = path[:len(path)-1]
			st = append(st, "(")
		}
	}
	bp(n)
	return res
}

// @lc code=end
