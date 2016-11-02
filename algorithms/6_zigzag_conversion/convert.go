package main

import "fmt"

func main() {
	fmt.Println(convert("dfkg", 3))
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	var ss = make([]string, numRows)

	for i := 0; i < len(s); i++ {
		var j int
		if i%(2*numRows-2) < numRows {
			j = i % (2*numRows - 2)
		} else {
			j = 2*numRows - 2 - i%(2*numRows-2)
		}
		ss[j] += string(s[i])
	}

	convertedstring := ""
	for _, v := range ss {
		convertedstring += v
	}
	return convertedstring
}
