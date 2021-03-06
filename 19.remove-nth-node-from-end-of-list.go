package main

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 *
 * https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (35.11%)
 * Likes:    3478
 * Dislikes: 234
 * Total Accepted:    644.3K
 * Total Submissions: 1.8M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * Given a linked list, remove the n-th node from the end of list and return
 * its head.
 *
 * Example:
 *
 *
 * Given linked list: 1->2->3->4->5, and n = 2.
 *
 * After removing the second node from the end, the linked list becomes
 * 1->2->3->5.
 *
 *
 * Note:
 *
 * Given n will always be valid.
 *
 * Follow up:
 *
 * Could you do this in one pass?
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}
	ns := make([]*ListNode, 0, 10)
	current := head
	i := 0

	for ; current != nil; i++ {
		ns = append(ns, current)
		current = current.Next
	}

	if n == len(ns) {
		head = head.Next
	} else if n == 1 {
		ns[i-n-1].Next = nil
	} else {
		ns[i-n-1].Next = ns[i-n+1]
	}

	return head
}

// @lc code=end
