package array

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	sorted := nums[:]
	numMap := make(map[int]int)
	sort.Ints(sorted)
	for i := len(sorted) - 1; i >= 0; i-- {
		numMap[sorted[i]] = i
	}
	res := make([]int, len(nums))
	for i, n := range nums {
		res[i] = numMap[n]
	}
	return res
}
