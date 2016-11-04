package main

import (
	"testing"
)

func Test_myAtoi(t *testing.T) {

	if myAtoi("") != 0 {
		t.Error("Test case 0 failed")
	}

	if myAtoi("-10") != -10 {
		t.Error("Test case 1 failed")
	}

	if myAtoi("99999999999999999999999999") != 2147483647 {
		t.Error("Test case 2 failed")
	}
}
