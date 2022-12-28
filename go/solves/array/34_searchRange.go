package array

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
进阶：你可以设计并实现时间复杂度为 $O(\log n)$ 的算法解决此问题吗？

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]
#
*/

func searchRange(nums []int, target int) []int {
	leftBound, rightBound := findLeft(nums, target), findRight(nums, target)
	// target 在左边
	if leftBound == -2 && rightBound == -2 {
		return []int{-1, -1}
	}
	// 在中间存在
	if rightBound-leftBound > 1 {
		return []int{leftBound + 1, rightBound - 1}
	}
	// 在中间，不存在
	return []int{-1, -1}
}

func findLeft(nums []int, target int) int {
	res := -2
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
			res = high
		}
	}
	return res
}

func findRight(nums []int, target int) int {
	res := -2
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
			res = low
		}
	}
	return res
}
