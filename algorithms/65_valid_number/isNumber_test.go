package main

import (
	"testing"
)

func Test_isNumber(t *testing.T) {

	if !isNumber("1.005") {
		t.Error("Test case 0 failed")
	}

	if !isNumber("1e5") {
		t.Error("Test case 1 failed")
	}

	if isNumber("1.xxx005") {
		t.Error("Test case 2 failed")
	}
}
