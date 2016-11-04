package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse(1991))
}

func reverse(x int) int {
	reverseNumber := 0
	for ; x != 0; x /= 10 {
		reverseNumber = reverseNumber*10 + x%10
		if reverseNumber > 2147483647 || reverseNumber < -2147483648 {
			return 0
		}
	}
	return reverseNumber
}
