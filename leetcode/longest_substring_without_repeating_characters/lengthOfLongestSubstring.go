package main

func lengthOfLongestSubstring(s string) int {
	var max int
	var pos [256]int
	for k, _ := range pos {
		pos[k] = -1
	}
	for i, j := 0, 0; i < len(s); i++ {
		c := int(s[i])
		if pos[c] >= j {
			j = pos[c] + 1
		} else {
		}
		pos[c] = i
		if 1 + i - j > max {
			max = i - j + 1
		}
	}
	return max
}
