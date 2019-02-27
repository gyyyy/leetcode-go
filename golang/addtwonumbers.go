package golang

import (
	"strconv"
)

// 2. Add Two Numbers
//
// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order and each of their nodes contain a single digit.
// Add the two numbers and return it as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Example:
// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() (s string) {
	for m := n; ; m = m.Next {
		s += strconv.Itoa(m.Val)
		if m.Next == nil {
			break
		}
		s += " -> "
	}
	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		v    = l1.Val + l2.Val
		node = &ListNode{v % 10, nil}
		pre  = node
	)
	v /= 10
	for n1, n2 := l1.Next, l2.Next; n1 != nil || n2 != nil || v > 0; v /= 10 {
		if n1 != nil {
			v += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			v += n2.Val
			n2 = n2.Next
		}
		n := &ListNode{v % 10, nil}
		pre.Next, pre = n, n
	}
	return node
}
