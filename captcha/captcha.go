package captcha

import "strconv"

func LeftOperand(pattern, left int) string {
	if pattern == 2 {
		return "One"
	}
	return strconv.Itoa(left)
}
