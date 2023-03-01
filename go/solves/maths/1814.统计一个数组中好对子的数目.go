package maths

/*
 * @lc app=leetcode.cn id=1814 lang=golang
 *
 * [1814] 统计一个数组中好对子的数目
 * nums[i] + rev(nums[j]) = nums[j] + rev[nums[i]]
 * ==>
 * nums[i] - rev(nums[i]) = nums[j] - rev[nums[j]]
 * 记录每个一个gap的次数x，下次再出现的话，就多了x对
 */

// @lc code=start
func countNicePairs(nums []int) int {
	m := make(map[int]int)
	res := 0
	mod := int(1e9 + 7)
	for i := 0; i < len(nums); i++ {
		cur := nums[i] - rev(nums[i])
		res = (res + m[cur]) % mod
		m[cur]++
	}
	return res
}

func rev(a int) int {
	res := 0
	for a > 0 {
		res = res*10 + a%10
		a = a / 10
	}
	return res
}

// @lc code=end
