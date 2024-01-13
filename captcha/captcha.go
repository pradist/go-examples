package captcha

import "strconv"

func LeftOperand(pattern, left int) string {
	if pattern == 2 {
		if left == 2 {
			return "Two"
		}
		return "One"
	}
	return strconv.Itoa(left)
}
