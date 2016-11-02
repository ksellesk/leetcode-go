package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abb"))

}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	lps := s[0:1]
	for i := 1; i < 2*len(s)-2; i++ {
		j, k := 0, 0
		if i%2 == 0 {
			j, k = i/2, i/2
		} else {
			j, k = (i+1)/2, i/2
		}

		for j > 0 && k < len(s)-1 {
			if s[j-1] != s[k+1] {
				break
			}
			j--
			k++
		}

		if k+1-j > len(lps) {
			lps = s[j : k+1]
		}
	}
	return lps
}
