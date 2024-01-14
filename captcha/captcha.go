package captcha

import "strconv"

func LeftOperand(pattern, left int) string {
	if pattern == 2 {
		if left == 2 {
			return "Two"
		}
		if left == 3 {
			return "Three"
		}
		if left == 4 {
			return "Four"
		}
		return "One"
	}
	return strconv.Itoa(left)
}
