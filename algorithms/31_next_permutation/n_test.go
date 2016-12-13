package main

import (
	"testing"
)

func Test_nextPermutation(t *testing.T) {

	nums := []int{1, 2}
	nums2 := []int{2, 1}
	nextPermutation(nums)
	for k, v := range nums {
		if nums2[k] != v {
			t.Error("Test case 0 failed")
		}
	}

}
