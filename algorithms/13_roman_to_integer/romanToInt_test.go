package main

import (
	"testing"
)

func Test_romanToInt(t *testing.T) {

	if romanToInt("DCXXI") != 621 {
		t.Error("Test case 0 failed")
	}

}
