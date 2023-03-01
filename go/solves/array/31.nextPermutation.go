package array

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				reverse(nums, i+1, len(nums)-1)
				return
			}
		}
	}
	// 到这里了说明整个数组都是倒序了，反转一下便可
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, from, to int) {
	for from < to {
		nums[from], nums[to] = nums[to], nums[from]
		from++
		to--
	}
}

// @lc code=end
