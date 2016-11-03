package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	var board [][]string
	var ans = make([]int, n)

	pos := 0
	count := 0
	for {
		hasPos := false
		for ; ans[pos] < n; ans[pos]++ {
			valid := true
			for i := 0; i < pos; i++ {
				if ans[i] == ans[pos] || ans[i]-ans[pos] == i-pos || ans[i]-ans[pos] == pos-i {
					valid = false
				}
			}
			if valid {
				hasPos = true
				break
			}
		}

		if !hasPos {
			pos--
		} else if pos == n-1 {
			board = append(board, formatAns(ans))
			count++
		} else {
			pos++
			ans[pos] = -1
		}

		if pos == -1 {
			break
		}
		ans[pos]++
	}
	return board
}

func formatAns(ans []int) []string {
	n := len(ans)
	var ss = make([]string, n)
	var s = make([]byte, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == ans[i] {
				s[j] = 'Q'
			} else {
				s[j] = '.'
			}
		}
		ss[i] = string(s)
	}

	return ss
}
