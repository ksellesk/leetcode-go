package main

import (
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	strs := []string{"xxx", "xxxa", "xxjdhldl"}
	if longestCommonPrefix(strs) != "xx" {
		t.Error("Test case 0 failed")
	}

}
