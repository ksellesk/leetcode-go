package main

import "fmt"

func main() {
	nums1, nums2 := []int{3}, []int{1, 2, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)

	if l1 == 0 && l2 == 0 {
		panic("Invalid input")
	}

	if l1 > l2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	if l1 == 0 {
		return float64(nums2[l2/2]+nums2[(l2-1)/2])/2.0
	}

	head, tail := 0, l1
	mid1, mid2 := 0, 0
	for head <= tail {
		mid1 = (head+tail)/2
		mid2 = (l1+l2+1)/2-mid1
		if mid1 < l1 &&  nums1[mid1] < nums2[mid2-1] {
			head = mid1 + 1
		} else if mid1 > 0 && nums1[mid1-1] > nums2[mid2] {
			tail = mid1 - 1
		} else {
			break
		}
	}

	var left, right int
	if mid1 == 0 {
		left = nums2[mid2-1]
	} else if mid2 == 0 {
		left = nums1[mid1-1]
	} else {
		if nums1[mid1-1] > nums2[mid2-1] {
			left = nums1[mid1-1]
		} else {
			left = nums2[mid2-1]
		}
	}

	if (l1 + l2) % 2 == 1 {
		return float64(left)
	}

	if mid1 == l1 {
		right = nums2[mid2]
	} else if mid2 == l2 {
		right = nums1[mid1]
	} else {
		if nums1[mid1] < nums2[mid2] {
			right = nums1[mid1]
		} else {
			right = nums2[mid2]
		}
	}

	return float64(left+right)/2.0
}
