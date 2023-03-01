package linklist

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	total := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		total++
	}
	k = k % total
	if k == 0 {
		return head
	}

	j := total - k
	cur = head
	for j > 1 {
		cur = cur.Next
		j--
	}
	last := cur
	first := cur.Next
	last.Next = nil
	cur = first
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = head
	return first
}

// @lc code=end
