package linklist

/*
 * @lc app=leetcode.cn id=1669 lang=golang
 *
 * [1669] 合并两个链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	list2Last := list2
	for list2Last.Next != nil {
		list2Last = list2Last.Next
	}
	list1Cut := list1
	list1Next := list1
	for a > 1 {
		list1Cut = list1Cut.Next
		a--
	}
	for b > -1 {
		list1Next = list1Next.Next
		b--
	}
	list1Cut.Next = list2
	list2Last.Next = list1Next
	return list1
}

// @lc code=end
