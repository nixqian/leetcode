package array

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	a := 0
	b := 1
	for a < len(nums) && b < len(nums) {
		if nums[a] == nums[b] {
			b++
		} else {
			nums[a+1], nums[b] = nums[b], nums[a+1]
			a++
			b++
		}
	}
	return a + 1
}

// @lc code=end
