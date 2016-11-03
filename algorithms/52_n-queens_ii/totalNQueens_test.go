package main

import (
	"testing"
)

func Test_totalNQueens(t *testing.T) {

	if totalNQueens(8) != 92 {
		t.Error("Test case 0 failed")
	}

}
