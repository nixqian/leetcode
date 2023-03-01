package greed

/*
 * @lc app=leetcode.cn id=2139 lang=golang
 *
 * [2139] 得到目标值的最少行动次数
 */

// @lc code=start
func minMoves(target int, maxDoubles int) int {
	cnt := 0
	for target > 1 {
		if maxDoubles == 0 {
			return cnt + target - 1
		}
		if target%2 == 1 {
			cnt++
			target--
		}
		maxDoubles--
		target = target / 2
		cnt++
	}
	return cnt
}

// @lc code=end
