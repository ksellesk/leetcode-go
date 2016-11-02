package main

import (
	"testing"
)

func Test_longestPalindrome(t *testing.T) {

	if longestPalindrome("abbcabba") != "abba" {
		t.Error("Test case 0 failed")
	}
}
