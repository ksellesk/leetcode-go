package main

func totalNQueens(n int) int {
	count := 0
	var ans = make([]int, n)
	pos := 0
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
	return count
}
