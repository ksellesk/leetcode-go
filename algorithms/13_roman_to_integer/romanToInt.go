package main

func romanToInt(s string) int {
	result := 0

	ints := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(strs); i++ {
		for len(s) >= len(strs[i]) && strs[i] == s[:len(strs[i])] {
			result += ints[i]
			s = s[len(strs[i]):]
		}
		if len(s) == 0 {
			break
		}
	}

	return result
}
