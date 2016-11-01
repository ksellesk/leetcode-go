package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(112211))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	counts := 1
	for x/counts >= 10 {
		counts = counts * 10
	}

	for x > 0 {
		left := x / counts
		right := x % 10

		if left != right {
			return false
		}

		x = (x % counts) / 10
		counts = counts / 100

	}
	return true
}
