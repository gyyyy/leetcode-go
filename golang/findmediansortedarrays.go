package golang

// 4. Median of Two Sorted Arrays
//
// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
// You may assume nums1 and nums2 cannot be both empty.
//
// Example 1:
// nums1 = [1, 3]
// nums2 = [2]
// The median is 2.0
//
// Example 2:
// nums1 = [1, 2]
// nums2 = [3, 4]
// The median is (2 + 3)/2 = 2.5
//
// Solution from LeetCode:
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	m, n := len(nums1), len(nums2)
// 	if m == 0 && n == 0 {
// 		return 0
// 	} else if m > n {
// 		nums1, nums2 = nums2, nums1
// 		m, n = n, m
// 	}
// 	for min, max, half := 0, m, (m+n+1)/2; min <= max; {
// 		var (
// 			i = (min + max) / 2
// 			j = half - i
// 		)
// 		if i < max && nums2[j-1] > nums1[i] {
// 			min = i + 1
// 		} else if i > min && nums1[i-1] > nums2[j] {
// 			max = i - 1
// 		} else {
// 			var left int
// 			if i == 0 {
// 				left = nums2[j-1]
// 			} else if j == 0 {
// 				left = nums1[i-1]
// 			} else if v1, v2 := nums1[i-1], nums2[j-1]; v1 >= v2 {
// 				left = v1
// 			} else {
// 				left = v2
// 			}
// 			if (m+n)%2 == 1 {
// 				return float64(left)
// 			}
// 			var right int
// 			if i == m {
// 				right = nums2[j]
// 			} else if j == n {
// 				right = nums1[i]
// 			} else if v1, v2 := nums1[i], nums2[j]; v1 <= v2 {
// 				right = v1
// 			} else {
// 				right = v2
// 			}
// 			return float64(left+right) / 2
// 		}
// 	}
// 	return 0
// }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		l1, l2 = len(nums1), len(nums2)
		m1, m2 = (l1 + l2) / 2, (l1 + l2) % 2
	)
	if l1 == 0 {
		if l2 == 0 {
			return 0
		}
		if m2 != 0 {
			return float64(nums2[m1])
		}
		return float64(nums2[m1-1]+nums2[m1]) / 2
	} else if l2 == 0 {
		if m2 != 0 {
			return float64(nums1[m1])
		}
		return float64(nums1[m1-1]+nums1[m1]) / 2
	}
	var pre, cur int
	for i, j, n := 0, 0, 0; n <= m1; n++ {
		if pre = cur; i >= l1 {
			cur = nums2[j]
			j++
		} else if j >= l2 {
			cur = nums1[i]
			i++
		} else if v1, v2 := nums1[i], nums2[j]; v1 <= v2 {
			cur = v1
			i++
		} else {
			cur = v2
			j++
		}
	}
	if m2 != 0 {
		return float64(cur)
	}
	return float64(pre+cur) / 2
}
