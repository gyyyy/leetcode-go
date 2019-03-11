package leetcode

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	if sum := twoSum([]int{2, 7, 11, 15}, 9); sum == nil || len(sum) != 2 {
		t.Fail()
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	if addTwoNumbers(l1, l2).String() != "7 -> 0 -> 8" {
		t.Fail()
	}
	l1 = &ListNode{5, nil}
	l2 = &ListNode{5, nil}
	if addTwoNumbers(l1, l2).String() != "0 -> 1" {
		t.Fail()
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	if lengthOfLongestSubstring("abcabcbb") != 3 ||
		lengthOfLongestSubstring("bbbbb") != 1 ||
		lengthOfLongestSubstring("pwwkew") != 3 ||
		lengthOfLongestSubstring("") != 0 ||
		lengthOfLongestSubstring(" ") != 1 ||
		lengthOfLongestSubstring("au") != 2 ||
		lengthOfLongestSubstring("abba") != 2 ||
		lengthOfLongestSubstring("aab") != 2 {
		t.Fail()
	}
}

func TestFindMedianSortedArrays(t *testing.T) {
	if findMedianSortedArrays([]int{1, 3}, []int{2, 4}) != 2.5 ||
		findMedianSortedArrays([]int{1, 3}, []int{2, 12, 13, 22, 31}) != 12 ||
		findMedianSortedArrays([]int{}, []int{}) != 0 ||
		findMedianSortedArrays([]int{1}, []int{}) != 1 ||
		findMedianSortedArrays([]int{}, []int{2}) != 2 {
		t.Fail()
	}
}
