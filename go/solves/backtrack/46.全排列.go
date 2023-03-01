package backtrack

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	var bp func()
	bp = func() {
		if len(path) == len(nums) {
			result = append(result, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			bp()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	bp()
	return result
}

// @lc code=end

/*
1. 递归函数
2. 结束条件
3. 单层搜索逻辑
*/
