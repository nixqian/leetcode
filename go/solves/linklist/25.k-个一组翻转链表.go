package linklist

import "fmt"

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	headK := head
	for headK != nil {
		tail := headK
		i := 1
		for ; i < k && tail != nil; i++ {
			tail = tail.Next
		}
		fmt.Println(i)
		if i != k {
			return dummy.Next
		}
		fmt.Println(cur.Val, headK.Val, tail.Val)
		idx := headK
		nextHeadK := tail.Next
		for idx != tail.Next {
			tmp := idx.Next
			idx.Next = cur.Next
			cur.Next = idx
			idx = tmp
		}
		cur = headK
		headK = nextHeadK
	}
	return dummy.Next
}

// @lc code=end
