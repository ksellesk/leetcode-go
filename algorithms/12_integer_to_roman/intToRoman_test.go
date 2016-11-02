package main

import (
	"testing"
)

func Test_intToRoman(t *testing.T) {

	if intToRoman(1) != "I" {
		t.Error("Test case 0 failed")
	}

	if intToRoman(1111) != "MCXI" {
		t.Error("Test case 1 failed")
	}

}
