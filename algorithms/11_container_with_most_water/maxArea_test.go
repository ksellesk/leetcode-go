package main

import "testing"

func Test_maxArea(t *testing.T) {
	height := []int{2, 3 ,4, 1}
	if maxArea(height) != 4 {
		t.Error("Test case 0 failed")
	}
}
