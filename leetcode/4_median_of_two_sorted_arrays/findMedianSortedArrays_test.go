package main

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	nums1, nums2 := []int{}, []int{1}
	if findMedianSortedArrays(nums1, nums2) != 1.0 {
		t.Error("Test case 0 failed")
	}
}
