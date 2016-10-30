package main

import (
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{2, 3, 6, 7}
	result := twoSum(nums, 13)
	if result[0] != 2 && result[1] != 3 {
		t.Error("Test case 0 failed")
	}

	if twoSum(nums, 14) != nil {
		t.Error("Test case 1 failed")
	}


}
