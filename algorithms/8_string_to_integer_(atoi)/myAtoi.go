package main

import (
	"fmt"
)

func main() {
	fmt.Println(myAtoi("9223372036854775809"))
}

func myAtoi(str string) int {
	if len(str) < 1 {
		return 0
	}

	i := 0
	for ; i < len(str) && str[i] == ' '; i++ {
	}

	neg := false
	if str[i] == '+' {
		i++
	} else if str[i] == '-' {
		neg = true
		i++
	}

	v := 0
	for j := 0; i < len(str); i++ {
		if str[i] > '9' || str[i] < '0' || j > 10 {
			break
		} else {
			v = v*10 + int(str[i]-48)
		}
		j++
	}
	if neg {
		v = 0 - v
	}

	if v > 2147483647 {
		return 2147483647
	}
	if v < -2147483648 {
		return -2147483648
	}
	return v
}
