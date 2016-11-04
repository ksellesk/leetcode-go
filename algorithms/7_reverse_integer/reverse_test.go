package main

import (
	"testing"
)

func Test_reverse(t *testing.T) {

	if reverse(6402) != 2046 {
		t.Error("Test case 0 failed")
	}

}
