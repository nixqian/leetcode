package linklist

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	newList := []*ListNode{}
	for _, list := range lists {
		if list != nil {
			newList = append(newList, list)
		}
	}
	head := &ListNode{}
	cur := head
	for len(newList) != 0 {
		minIdx := -1
		minVal := math.MaxInt
		for idx, list := range newList {
			if list.Val < minVal {
				minVal = list.Val
				minIdx = idx
			}
		}
		cur.Next = newList[minIdx]
		cur = cur.Next
		newList[minIdx] = newList[minIdx].Next
		if newList[minIdx] == nil { // 删除用完的list
			newList = append(newList[:minIdx], newList[minIdx+1:]...)
		}
	}
	return head.Next
}

// @lc code=end
