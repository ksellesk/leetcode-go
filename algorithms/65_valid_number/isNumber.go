package main

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	space := false
	exp := false
	dot := false
	number := false
	neg := false

	for ; i < len(s); i++ {
		c := s[i]
		if c == ' ' {
			space = true
		} else if space == true {
			return false
		} else if (c == 'e' || c == 'E') && exp == false && number == true {
			exp = true
			number = false
			dot = true
			neg = false
		} else if c == '.' && dot == false {
			dot = true
			neg = true
		} else if c >= '0' && c <= '9' {
			number = true
		} else if (c == '+' || c == '-') && neg == false && number == false {
			neg = true
		} else {
			return false
		}
	}
	return number
}
