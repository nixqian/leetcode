package backtrack

/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) > 1 {
			res = append(res, append([]int{}, path...))
		}
		used := map[int]struct{}{}
		for i := start; i < len(nums); i++ {
			if len(path) != 0 && nums[i] < path[len(path)-1] {
				continue
			}
			if _, ok := used[nums[i]]; ok {
				continue
			}
			used[nums[i]] = struct{}{}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

// @lc code=end
