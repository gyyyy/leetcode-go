package golang

// 3. Longest Substring Without Repeating Characters
//
// Given a string, find the length of the longest substring without repeating characters.
//
// Example 1:
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
//
// Example 2:
// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
// Example 3:
// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
//              Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) (n int) {
	l := len(s)
	if l < 2 {
		return l
	}
	m := make(map[byte]int)
	for i, j := 0, 0; j < l; j++ {
		c := s[j]
		if v, ok := m[c]; ok && v > i {
			i = v
		}
		if k := j - i + 1; k > n {
			n = k
		}
		m[c] = j + 1
	}
	return
}
