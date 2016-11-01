package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	if !isPalindrome(112211) {
		t.Error("Test case 0 failed")
	}
}
