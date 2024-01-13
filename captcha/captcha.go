package captcha

import "strconv"

func LeftOperand(pattern, left int) string {
	if left == 2 {
		return strconv.Itoa(left)
	}
	if left == 3 {
		return strconv.Itoa(left)
	}
	if left == 4 {
		return strconv.Itoa(left)
	}
	return strconv.Itoa(left)
}
