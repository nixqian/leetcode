package array

/*
 * @lc app=leetcode.cn id=2070 lang=golang
 *
 * [2070] 每一个查询的最大美丽值
 */

// @lc code=start
func maximumBeauty(items [][]int, queries []int) []int {
	ans := []int{}
	for _, query := range queries {
		maxV := 0
		for _, item := range items {
			if item[0] <= query && maxV < item[1] {
				maxV = item[1]
			}
		}
		ans = append(ans, maxV)
	}
	return ans
}

// @lc code=end
