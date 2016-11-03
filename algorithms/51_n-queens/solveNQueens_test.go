package main

import (
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	expect := [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}
	for k1, v1 := range solveNQueens(4) {
		for k2, v2 := range v1 {
			if expect[k1][k2] != v2 {
				t.Error("Test case 0 failed")
				println(v2)
			}
		}
	}
}
