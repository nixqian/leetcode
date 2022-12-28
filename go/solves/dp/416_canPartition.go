package dp

/*
题目难易：中等

给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意: 每个数组中的元素不会超过 100 数组的大小不会超过 200

示例 1: 输入: [1, 5, 11, 5] 输出: true 解释: 数组可以分割成 [1, 5, 5] 和 [11].

示例 2: 输入: [1, 2, 3, 5] 输出: false 解释: 数组不能分割成两个元素和相等的子集.

提示：

1 <= nums.length <= 200
1 <= nums[i] <= 100
#
*/

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 { // 奇数和无法分割
		return false
	}
	var intMax = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	target := sum / 2
	dp := make([]int, target+1)
	// dp[j] = max(dp[j], dp[j - nums[i]] + nums[i]);
	for i := 1; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- { // 遍历背包容量
			dp[j] = intMax(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}
