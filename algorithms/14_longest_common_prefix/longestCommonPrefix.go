package main

import "fmt"

func main() {
	strs := []string{"aa", "a"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	commonPrefix := ""
	for i := 0; i < len(strs[0]); i++ {
		j := 1

		for ; j < len(strs); j++ {
			if len(strs[j]) < i+1 || strs[j][i] != strs[0][i] {
				break
			}
		}

		if j < len(strs) {
			break
		} else {
			commonPrefix = commonPrefix + string(strs[0][i])
		}
	}

	return commonPrefix
}
