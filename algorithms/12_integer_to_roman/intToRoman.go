package main

func intToRoman(num int) string {
	result := ""

	ints := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(ints); i++ {
		for ; num >= ints[i]; num -= ints[i] {
			result = result + strs[i]
		}
		if num == 0 {
			break
		}
	}
	return result
}
