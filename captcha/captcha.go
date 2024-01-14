package captcha

import "strconv"

var number = []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

func LeftOperand(pattern, left int) string {
	if pattern == 2 {
		if left == 2 {
			return number[left-1]
		}
		if left == 3 {
			return number[left-1]
		}
		if left == 4 {
			return number[left-1]
		}
		return number[left-1]
	}
	return strconv.Itoa(left)
}
