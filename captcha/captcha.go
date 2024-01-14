package captcha

import "strconv"

var number = []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

func LeftOperand(pattern, left int) string {
	if pattern == 2 {
		return number[left]
	}
	return strconv.Itoa(left)
}
