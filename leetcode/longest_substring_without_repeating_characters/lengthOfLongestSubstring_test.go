package main

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	s := "aflhgalgfafgahagabcdefghijklmnopqrstuvwxyz"
	if lengthOfLongestSubstring(s) != 26 {
		t.Error("Test case 0 failed")
	}
}
