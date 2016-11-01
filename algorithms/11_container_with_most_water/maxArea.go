package main

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	head, tail := 0, len(height)-1
	h1, h2 := height[head], height[tail]
	h, flag := min(h1, h2)
	area := (tail - head) * h
	max := area

	for head < tail-1 {
		if flag {
			head += 1
			if height[head] > h1 {
				h1 = height[head]
			} else {
				continue
			}
		} else {
			tail -= 1
			if height[tail] > h2 {
				h2 = height[tail]
			} else {
				continue
			}
		}

		h, flag = min(h1, h2)
		area = (tail - head) * h
		if area > max {
			max = area
		}

	}

	return max
}

func min(x, y int) (int, bool) {
	if x < y {
		return x, true
	}
	return y, false
}
